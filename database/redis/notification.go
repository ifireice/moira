package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"

	"github.com/moira-alert/moira"
	"github.com/moira-alert/moira/database/redis/reply"
)

// GetNotifications gets ScheduledNotifications in given range and full range
func (connector *DbConnector) GetNotifications(start, end int64) ([]*moira.ScheduledNotification, int64, error) {
	c := connector.pool.Get()
	defer c.Close()

	c.Send("MULTI")
	c.Send("ZRANGE", notifierNotificationsKey, start, end)
	c.Send("ZCARD", notifierNotificationsKey)
	rawResponse, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to EXEC: %s", err.Error())
	}
	if len(rawResponse) == 0 {
		return make([]*moira.ScheduledNotification, 0), 0, nil
	}
	total, err := redis.Int64(rawResponse[1], nil)
	if err != nil {
		return nil, 0, err
	}
	notifications, err := reply.Notifications(rawResponse[0], nil)
	if err != nil {
		return nil, 0, err
	}
	return notifications, total, nil
}

// RemoveAllNotifications delete all notifications
func (connector *DbConnector) RemoveAllNotifications() error {
	c := connector.pool.Get()
	defer c.Close()

	if _, err := c.Do("DEL", notifierNotificationsKey); err != nil {
		return fmt.Errorf("failed to remove %s: %s", notifierNotificationsKey, err.Error())
	}

	return nil
}

// RemoveNotification delete notifications by key = timestamp + contactID + subID
func (connector *DbConnector) RemoveNotification(notificationKey string) (int64, error) {
	notifications, _, err := connector.GetNotifications(0, -1)
	if err != nil {
		return 0, err
	}

	foundNotifications := make([]*moira.ScheduledNotification, 0)
	for _, notification := range notifications {
		timestamp := strconv.FormatInt(notification.Timestamp, 10)
		contactID := notification.Contact.ID
		subID := moira.UseString(notification.Event.SubscriptionID)
		idstr := strings.Join([]string{timestamp, contactID, subID}, "")
		if idstr == notificationKey {
			foundNotifications = append(foundNotifications, notification)
		}
	}
	return connector.removeNotifications(foundNotifications)
}

func (connector *DbConnector) removeNotifications(notifications []*moira.ScheduledNotification) (int64, error) {
	if len(notifications) == 0 {
		return 0, nil
	}

	c := connector.pool.Get()
	defer c.Close()

	c.Send("MULTI")
	for _, notification := range notifications {
		notificationString, err := json.Marshal(notification)
		if err != nil {
			return 0, err
		}
		c.Send("ZREM", notifierNotificationsKey, notificationString)
	}
	response, err := redis.Ints(c.Do("EXEC"))
	if err != nil {
		return 0, fmt.Errorf("failed to remove notifier-notification: %s", err.Error())
	}
	total := 0
	for _, val := range response {
		total += val
	}
	return int64(total), nil
}

// FetchNotifications fetch notifications by given timestamp and delete it
func (connector *DbConnector) FetchNotifications(to int64, limit int64) ([]*moira.ScheduledNotification, error){
	// No limit
	if limit == 0 {
		return connector.fetchNotificationsNoLimit(to)
	}

	count, err := connector.notificationsCount(to)
	if err != nil {
		return nil, err
	}

	// Hope count will be not greater then limit when we call fetchNotificationsNoLimit
	if *count < limit / 2 {
		return connector.fetchNotificationsNoLimit(to)
	}

	return connector.fetchNotificationsWithLimit(to, limit)
}

func (connector *DbConnector) notificationsCount(to int64) (*int64, error){
	c := connector.pool.Get()
	defer c.Close()

	count, err := redis.Int64(c.Do("ZCOUNT", notifierNotificationsKey, "-inf", to))
	if err != nil {
		return nil, fmt.Errorf("failed to ZCOUNT to notificationsCount: %s", err)
	}
	return &count, nil
}

// Drops all notifivtsions with max timestamp
func limitNotifications(notifications []*moira.ScheduledNotification) ([]*moira.ScheduledNotification) {
	last_ts := notifications[len(notifications) - 1].Timestamp
	i := len(notifications) - 1
	for ;i >= 0; i-- {
		if notifications[i].Timestamp != last_ts {
			break
		}
	}

	if i == -1 {
		return notifications;
	}

	return notifications[:i+1]
}

// fetchNotificationsWithLimit reads and drops notifications from DB with limit
func (connector *DbConnector) fetchNotificationsWithLimit(to int64, limit int64) ([]*moira.ScheduledNotification, error) {
	// т.к. при удалении испольузется Watch нотификации могут не удалится из-за изменения данных в другом потоке
	// см. https://redis.io/topics/transactions
	for i := 0; i <= 10; i++ {
		res, err := connector.fetchNotificationsWithLimitDo(to, limit)
		if err == nil {
			return res, nil
		}

		if err.Error() != "Transaction error" {
			return nil, err
		}
	}
	return nil, fmt.Errorf("Transaction tries limit exided!")
}

// same as fetchNotificationsWithLimit, but only once
func (connector *DbConnector) fetchNotificationsWithLimitDo(to int64, limit int64) ([]*moira.ScheduledNotification, error) {
	c := connector.pool.Get()
	defer c.Close()

	c.Send("WATCH", notifierNotificationsKey)
	response, err := redis.Values(c.Do("ZRANGEBYSCORE", notifierNotificationsKey, "-inf", to, "LIMIT", 0, limit))
	if err != nil {
		return nil, fmt.Errorf("failed to EXEC: %s", err)
	}

	if len(response) == 0 {
		return make([]*moira.ScheduledNotification, 0), nil
	}

	notifications, err := reply.Notifications(response, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to EXEC: %s", err)
	}

	notificationsLimited := limitNotifications(notifications)
	last_ts := notificationsLimited[len(notificationsLimited) - 1].Timestamp
	if (len(notifications) == len(notificationsLimited)) {
		c.Send("UNWATCH")
		return connector.fetchNotificationsNoLimit(last_ts)
	}

	c.Send("MULTI")
	c.Send("ZREMRANGEBYSCORE", notifierNotificationsKey, "-inf", last_ts)
	deleted_count, err_delete := redis.Values(c.Do("EXEC"))
	if err_delete != nil {
		return nil, fmt.Errorf("failed to EXEC: %s", err_delete)
	}

	if deleted_count[0] == 0 {
		return nil, fmt.Errorf("Transaction error")
	}

	return notificationsLimited, nil
}

// FetchNotifications fetch notifications by given timestamp and delete it
func (connector *DbConnector) fetchNotificationsNoLimit(to int64) ([]*moira.ScheduledNotification, error) {
	c := connector.pool.Get()
	defer c.Close()

	c.Send("MULTI")
	c.Send("ZRANGEBYSCORE", notifierNotificationsKey, "-inf", to)
	c.Send("ZREMRANGEBYSCORE", notifierNotificationsKey, "-inf", to)
	response, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		return nil, fmt.Errorf("failed to EXEC: %s", err)
	}
	if len(response) == 0 {
		return make([]*moira.ScheduledNotification, 0), nil
	}
	return reply.Notifications(response[0], nil)
}

// AddNotification store notification at given timestamp
func (connector *DbConnector) AddNotification(notification *moira.ScheduledNotification) error {
	bytes, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	c := connector.pool.Get()
	defer c.Close()

	_, err = c.Do("ZADD", notifierNotificationsKey, notification.Timestamp, bytes)
	if err != nil {
		return fmt.Errorf("failed to add scheduled notification: %s, error: %s", string(bytes), err.Error())
	}
	return err
}

// AddNotifications store notification at given timestamp
func (connector *DbConnector) AddNotifications(notifications []*moira.ScheduledNotification, timestamp int64) error {
	c := connector.pool.Get()
	defer c.Close()

	c.Send("MULTI")
	for _, notification := range notifications {
		bytes, err := json.Marshal(notification)
		if err != nil {
			return err
		}
		c.Send("ZADD", notifierNotificationsKey, timestamp, bytes)
	}
	_, err := c.Do("EXEC")
	if err != nil {
		return fmt.Errorf("failed to EXEC: %s", err.Error())
	}
	return nil
}

var notifierNotificationsKey = "moira-notifier-notifications"
