package redis

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/op/go-logging"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/moira-alert/moira"
)

func TestScheduledNotification(t *testing.T) {
	logger, _ := logging.GetLogger("dataBase")
	dataBase := newTestDatabase(logger, config)
	dataBase.flush()
	defer dataBase.flush()

	Convey("ScheduledNotification manipulation", t, func() {
		now := time.Now().Unix()
		notificationNew := moira.ScheduledNotification{
			SendFail:  1,
			Timestamp: now + 3600,
		}
		notification := moira.ScheduledNotification{
			SendFail:  2,
			Timestamp: now,
		}
		notificationOld := moira.ScheduledNotification{
			SendFail:  3,
			Timestamp: now - 3600,
		}

		Convey("Test add and get by pages", func() {
			addNotifications(dataBase, []moira.ScheduledNotification{notification, notificationNew, notificationOld})
			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 3)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notificationOld, &notification, &notificationNew})

			actual, total, err = dataBase.GetNotifications(0, 0)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 3)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notificationOld})

			actual, total, err = dataBase.GetNotifications(1, 2)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 3)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification, &notificationNew})
		})

		Convey("Test fetch notifications", func() {
			actual, err := dataBase.FetchNotifications(now - 3600, 0)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notificationOld})

			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 2)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification, &notificationNew})

			actual, err = dataBase.FetchNotifications(now + 3600, 0)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification, &notificationNew})

			actual, total, err = dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 0)
			So(actual, ShouldResemble, make([]*moira.ScheduledNotification, 0))
		})

		Convey("Test remove notifications by key", func() {
			id1 := "id1"
			notification1 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  1,
				Timestamp: now,
			}
			notification2 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  2,
				Timestamp: now,
			}
			notification3 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  3,
				Timestamp: now + 3600,
			}
			addNotifications(dataBase, []moira.ScheduledNotification{notification1, notification2, notification3})
			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 3)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification1, &notification2, &notification3})

			total, err = dataBase.RemoveNotification(strings.Join([]string{fmt.Sprintf("%v", now), id1, id1}, ""))
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 2)

			actual, total, err = dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 1)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification3})

			total, err = dataBase.RemoveNotification(strings.Join([]string{fmt.Sprintf("%v", now+3600), id1, id1}, ""))
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 1)

			actual, total, err = dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 0)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{})

			actual, err = dataBase.FetchNotifications(now + 3600, 0)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{})
		})

		Convey("Test remove all notifications", func() {
			id1 := "id1"
			notification1 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  1,
				Timestamp: now,
			}
			notification2 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  2,
				Timestamp: now,
			}
			notification3 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  3,
				Timestamp: now + 3600,
			}
			addNotifications(dataBase, []moira.ScheduledNotification{notification1, notification2, notification3})
			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 3)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification1, &notification2, &notification3})

			err = dataBase.RemoveAllNotifications()
			So(err, ShouldBeNil)

			actual, total, err = dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 0)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{})

			actual, err = dataBase.FetchNotifications(now + 3600, 0)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{})
		})

		Convey("Test fetch notifications with limit if all notifications has diff timestamp", func() {
			addNotifications(dataBase, []moira.ScheduledNotification{notification, notificationNew, notificationOld})
			actual, err := dataBase.FetchNotifications(now + 6000, 1)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notificationOld})

			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 2)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification, &notificationNew})

			err = dataBase.RemoveAllNotifications()
			So(err, ShouldBeNil)
		})
		Convey("Test fetch notifications with limit little bit greater than count if all notifications has diff timestamp", func() {
			addNotifications(dataBase, []moira.ScheduledNotification{notification, notificationNew, notificationOld})
			actual, err := dataBase.FetchNotifications(now + 6000, 4)
			So(err, ShouldBeNil)

			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notificationOld, &notification})

			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 1)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notificationNew})

			err = dataBase.RemoveAllNotifications()
			So(err, ShouldBeNil)
		})
		Convey("Test fetch notifications with limit greater than count if all notifications has diff timestamp", func() {
			addNotifications(dataBase, []moira.ScheduledNotification{notification, notificationNew, notificationOld})
			actual, err := dataBase.FetchNotifications(now + 6000, 8)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notificationOld, &notification, &notificationNew})

			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 0)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{})

			err = dataBase.RemoveAllNotifications()
			So(err, ShouldBeNil)
		})
		Convey("Test fetch notifications with limit if all notifications has same timestamp", func() {
			now := time.Now().Unix()
			id1 := "id1"
			notification1 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  1,
				Timestamp: now,
			}
			notification2 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  2,
				Timestamp: now,
			}
			notification3 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  3,
				Timestamp: now,
			}
			addNotifications(dataBase, []moira.ScheduledNotification{notification1, notification2, notification3})

			actual, err := dataBase.FetchNotifications(now + 6000, 1)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification1, &notification2, &notification3})

			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 0)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{})

			err = dataBase.RemoveAllNotifications()
			So(err, ShouldBeNil)
		})
		Convey("Test fetch notifications with limit if two notifications has same ts and one grater ts", func() {
			now := time.Now().Unix()
			id1 := "id1"
			notification1 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  1,
				Timestamp: now,
			}
			notification2 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  2,
				Timestamp: now,
			}
			notification3 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  3,
				Timestamp: now + 3600,
			}
			addNotifications(dataBase, []moira.ScheduledNotification{notification1, notification2, notification3})

			actual, err := dataBase.FetchNotifications(now + 6000, 1)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification1, &notification2})

			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 1)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification3})

			err = dataBase.RemoveAllNotifications()
			So(err, ShouldBeNil)
		})
		Convey("Test fetch notifications with limit if one notifications has less ts and other same grater ts", func() {
			now := time.Now().Unix()
			id1 := "id1"
			notification1 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  1,
				Timestamp: now,
			}
			notification2 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  2,
				Timestamp: now + 3600,
			}
			notification3 := moira.ScheduledNotification{
				Contact:   moira.ContactData{ID: id1},
				Event:     moira.NotificationEvent{SubscriptionID: &id1},
				SendFail:  3,
				Timestamp: now + 3600,
			}
			addNotifications(dataBase, []moira.ScheduledNotification{notification1, notification2, notification3})

			actual, err := dataBase.FetchNotifications(now + 6000, 1)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification1})

			actual, total, err := dataBase.GetNotifications(0, -1)
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 2)
			So(actual, ShouldResemble, []*moira.ScheduledNotification{&notification2, &notification3})

			err = dataBase.RemoveAllNotifications()
			So(err, ShouldBeNil)
		})
	})
}

func addNotifications(dataBase moira.Database, notifications []moira.ScheduledNotification) {
	for _, notification := range notifications {
		err := dataBase.AddNotification(&notification)
		So(err, ShouldBeNil)
	}
}

func TestScheduledNotificationErrorConnection(t *testing.T) {
	logger, _ := logging.GetLogger("dataBase")
	dataBase := newTestDatabase(logger, emptyConfig)
	dataBase.flush()
	defer dataBase.flush()

	Convey("Should throw error when no connection", t, func() {
		actual1, total, err := dataBase.GetNotifications(0, 1)
		So(actual1, ShouldBeNil)
		So(total, ShouldEqual, 0)
		So(err, ShouldNotBeNil)

		total, err = dataBase.RemoveNotification("123")
		So(err, ShouldNotBeNil)
		So(total, ShouldEqual, 0)

		actual2, err := dataBase.FetchNotifications(0, 0)
		So(err, ShouldNotBeNil)
		So(actual2, ShouldBeNil)

		notification := moira.ScheduledNotification{}
		err = dataBase.AddNotification(&notification)
		So(err, ShouldNotBeNil)

		err = dataBase.AddNotifications([]*moira.ScheduledNotification{&notification}, 0)
		So(err, ShouldNotBeNil)

		err = dataBase.RemoveAllNotifications()
		So(err, ShouldNotBeNil)
	})
}
