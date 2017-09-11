// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira (interfaces: Database)

package mock_moira_alert

import (
	gomock "github.com/golang/mock/gomock"
	moira_alert "github.com/moira-alert/moira"
	tomb_v2 "gopkg.in/tomb.v2"
	time "time"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return _m.recorder
}

// AcquireTriggerCheckLock mocks base method
func (_m *MockDatabase) AcquireTriggerCheckLock(_param0 string, _param1 int) error {
	ret := _m.ctrl.Call(_m, "AcquireTriggerCheckLock", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcquireTriggerCheckLock indicates an expected call of AcquireTriggerCheckLock
func (_mr *MockDatabaseMockRecorder) AcquireTriggerCheckLock(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcquireTriggerCheckLock", arg0, arg1)
}

// AddNotification mocks base method
func (_m *MockDatabase) AddNotification(_param0 *moira_alert.ScheduledNotification) error {
	ret := _m.ctrl.Call(_m, "AddNotification", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNotification indicates an expected call of AddNotification
func (_mr *MockDatabaseMockRecorder) AddNotification(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddNotification", arg0)
}

// AddNotifications mocks base method
func (_m *MockDatabase) AddNotifications(_param0 []*moira_alert.ScheduledNotification, _param1 int64) error {
	ret := _m.ctrl.Call(_m, "AddNotifications", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNotifications indicates an expected call of AddNotifications
func (_mr *MockDatabaseMockRecorder) AddNotifications(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddNotifications", arg0, arg1)
}

// AddPatternMetric mocks base method
func (_m *MockDatabase) AddPatternMetric(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AddPatternMetric", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPatternMetric indicates an expected call of AddPatternMetric
func (_mr *MockDatabaseMockRecorder) AddPatternMetric(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddPatternMetric", arg0, arg1)
}

// DeleteTriggerCheckLock mocks base method
func (_m *MockDatabase) DeleteTriggerCheckLock(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTriggerCheckLock", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTriggerCheckLock indicates an expected call of DeleteTriggerCheckLock
func (_mr *MockDatabaseMockRecorder) DeleteTriggerCheckLock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTriggerCheckLock", arg0)
}

// DeleteTriggerThrottling mocks base method
func (_m *MockDatabase) DeleteTriggerThrottling(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTriggerThrottling", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTriggerThrottling indicates an expected call of DeleteTriggerThrottling
func (_mr *MockDatabaseMockRecorder) DeleteTriggerThrottling(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTriggerThrottling", arg0)
}

// DeregisterBot mocks base method
func (_m *MockDatabase) DeregisterBot(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeregisterBot", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterBot indicates an expected call of DeregisterBot
func (_mr *MockDatabaseMockRecorder) DeregisterBot(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterBot", arg0)
}

// DeregisterBots mocks base method
func (_m *MockDatabase) DeregisterBots() {
	_m.ctrl.Call(_m, "DeregisterBots")
}

// DeregisterBots indicates an expected call of DeregisterBots
func (_mr *MockDatabaseMockRecorder) DeregisterBots() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterBots")
}

// FetchNotificationEvent mocks base method
func (_m *MockDatabase) FetchNotificationEvent() (moira_alert.NotificationEvent, error) {
	ret := _m.ctrl.Call(_m, "FetchNotificationEvent")
	ret0, _ := ret[0].(moira_alert.NotificationEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNotificationEvent indicates an expected call of FetchNotificationEvent
func (_mr *MockDatabaseMockRecorder) FetchNotificationEvent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchNotificationEvent")
}

// FetchNotifications mocks base method
func (_m *MockDatabase) FetchNotifications(_param0 int64) ([]*moira_alert.ScheduledNotification, error) {
	ret := _m.ctrl.Call(_m, "FetchNotifications", _param0)
	ret0, _ := ret[0].([]*moira_alert.ScheduledNotification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNotifications indicates an expected call of FetchNotifications
func (_mr *MockDatabaseMockRecorder) FetchNotifications(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchNotifications", arg0)
}

// GetAllContacts mocks base method
func (_m *MockDatabase) GetAllContacts() ([]*moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetAllContacts")
	ret0, _ := ret[0].([]*moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllContacts indicates an expected call of GetAllContacts
func (_mr *MockDatabaseMockRecorder) GetAllContacts() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllContacts")
}

// GetChecksUpdatesCount mocks base method
func (_m *MockDatabase) GetChecksUpdatesCount() (int64, error) {
	ret := _m.ctrl.Call(_m, "GetChecksUpdatesCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecksUpdatesCount indicates an expected call of GetChecksUpdatesCount
func (_mr *MockDatabaseMockRecorder) GetChecksUpdatesCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetChecksUpdatesCount")
}

// GetContact mocks base method
func (_m *MockDatabase) GetContact(_param0 string) (moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetContact", _param0)
	ret0, _ := ret[0].(moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContact indicates an expected call of GetContact
func (_mr *MockDatabaseMockRecorder) GetContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContact", arg0)
}

// GetContacts mocks base method
func (_m *MockDatabase) GetContacts(_param0 []string) ([]*moira_alert.ContactData, error) {
	ret := _m.ctrl.Call(_m, "GetContacts", _param0)
	ret0, _ := ret[0].([]*moira_alert.ContactData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContacts indicates an expected call of GetContacts
func (_mr *MockDatabaseMockRecorder) GetContacts(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContacts", arg0)
}

// GetIDByUsername mocks base method
func (_m *MockDatabase) GetIDByUsername(_param0 string, _param1 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetIDByUsername", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDByUsername indicates an expected call of GetIDByUsername
func (_mr *MockDatabaseMockRecorder) GetIDByUsername(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetIDByUsername", arg0, arg1)
}

// GetMetricRetention mocks base method
func (_m *MockDatabase) GetMetricRetention(_param0 string) (int64, error) {
	ret := _m.ctrl.Call(_m, "GetMetricRetention", _param0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricRetention indicates an expected call of GetMetricRetention
func (_mr *MockDatabaseMockRecorder) GetMetricRetention(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricRetention", arg0)
}

// GetMetricsUpdatesCount mocks base method
func (_m *MockDatabase) GetMetricsUpdatesCount() (int64, error) {
	ret := _m.ctrl.Call(_m, "GetMetricsUpdatesCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricsUpdatesCount indicates an expected call of GetMetricsUpdatesCount
func (_mr *MockDatabaseMockRecorder) GetMetricsUpdatesCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricsUpdatesCount")
}

// GetMetricsValues mocks base method
func (_m *MockDatabase) GetMetricsValues(_param0 []string, _param1 int64, _param2 int64) (map[string][]*moira_alert.MetricValue, error) {
	ret := _m.ctrl.Call(_m, "GetMetricsValues", _param0, _param1, _param2)
	ret0, _ := ret[0].(map[string][]*moira_alert.MetricValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricsValues indicates an expected call of GetMetricsValues
func (_mr *MockDatabaseMockRecorder) GetMetricsValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricsValues", arg0, arg1, arg2)
}

// GetNotificationEventCount mocks base method
func (_m *MockDatabase) GetNotificationEventCount(_param0 string, _param1 int64) int64 {
	ret := _m.ctrl.Call(_m, "GetNotificationEventCount", _param0, _param1)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNotificationEventCount indicates an expected call of GetNotificationEventCount
func (_mr *MockDatabaseMockRecorder) GetNotificationEventCount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotificationEventCount", arg0, arg1)
}

// GetNotificationEvents mocks base method
func (_m *MockDatabase) GetNotificationEvents(_param0 string, _param1 int64, _param2 int64) ([]*moira_alert.NotificationEvent, error) {
	ret := _m.ctrl.Call(_m, "GetNotificationEvents", _param0, _param1, _param2)
	ret0, _ := ret[0].([]*moira_alert.NotificationEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotificationEvents indicates an expected call of GetNotificationEvents
func (_mr *MockDatabaseMockRecorder) GetNotificationEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotificationEvents", arg0, arg1, arg2)
}

// GetNotifications mocks base method
func (_m *MockDatabase) GetNotifications(_param0 int64, _param1 int64) ([]*moira_alert.ScheduledNotification, int64, error) {
	ret := _m.ctrl.Call(_m, "GetNotifications", _param0, _param1)
	ret0, _ := ret[0].([]*moira_alert.ScheduledNotification)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNotifications indicates an expected call of GetNotifications
func (_mr *MockDatabaseMockRecorder) GetNotifications(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNotifications", arg0, arg1)
}

// GetPatternMetrics mocks base method
func (_m *MockDatabase) GetPatternMetrics(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetPatternMetrics", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatternMetrics indicates an expected call of GetPatternMetrics
func (_mr *MockDatabaseMockRecorder) GetPatternMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPatternMetrics", arg0)
}

// GetPatternTriggerIDs mocks base method
func (_m *MockDatabase) GetPatternTriggerIDs(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetPatternTriggerIDs", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatternTriggerIDs indicates an expected call of GetPatternTriggerIDs
func (_mr *MockDatabaseMockRecorder) GetPatternTriggerIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPatternTriggerIDs", arg0)
}

// GetPatterns mocks base method
func (_m *MockDatabase) GetPatterns() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetPatterns")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPatterns indicates an expected call of GetPatterns
func (_mr *MockDatabaseMockRecorder) GetPatterns() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPatterns")
}

// GetSubscription mocks base method
func (_m *MockDatabase) GetSubscription(_param0 string) (moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetSubscription", _param0)
	ret0, _ := ret[0].(moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscription indicates an expected call of GetSubscription
func (_mr *MockDatabaseMockRecorder) GetSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscription", arg0)
}

// GetSubscriptions mocks base method
func (_m *MockDatabase) GetSubscriptions(_param0 []string) ([]*moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetSubscriptions", _param0)
	ret0, _ := ret[0].([]*moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions
func (_mr *MockDatabaseMockRecorder) GetSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscriptions", arg0)
}

// GetTagNames mocks base method
func (_m *MockDatabase) GetTagNames() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTagNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagNames indicates an expected call of GetTagNames
func (_mr *MockDatabaseMockRecorder) GetTagNames() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagNames")
}

// GetTagTriggerIDs mocks base method
func (_m *MockDatabase) GetTagTriggerIDs(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTagTriggerIDs", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagTriggerIDs indicates an expected call of GetTagTriggerIDs
func (_mr *MockDatabaseMockRecorder) GetTagTriggerIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagTriggerIDs", arg0)
}

// GetTagsSubscriptions mocks base method
func (_m *MockDatabase) GetTagsSubscriptions(_param0 []string) ([]*moira_alert.SubscriptionData, error) {
	ret := _m.ctrl.Call(_m, "GetTagsSubscriptions", _param0)
	ret0, _ := ret[0].([]*moira_alert.SubscriptionData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsSubscriptions indicates an expected call of GetTagsSubscriptions
func (_mr *MockDatabaseMockRecorder) GetTagsSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTagsSubscriptions", arg0)
}

// GetTrigger mocks base method
func (_m *MockDatabase) GetTrigger(_param0 string) (moira_alert.Trigger, error) {
	ret := _m.ctrl.Call(_m, "GetTrigger", _param0)
	ret0, _ := ret[0].(moira_alert.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrigger indicates an expected call of GetTrigger
func (_mr *MockDatabaseMockRecorder) GetTrigger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTrigger", arg0)
}

// GetTriggerCheckIDs mocks base method
func (_m *MockDatabase) GetTriggerCheckIDs(_param0 []string, _param1 bool) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerCheckIDs", _param0, _param1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerCheckIDs indicates an expected call of GetTriggerCheckIDs
func (_mr *MockDatabaseMockRecorder) GetTriggerCheckIDs(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerCheckIDs", arg0, arg1)
}

// GetTriggerChecks mocks base method
func (_m *MockDatabase) GetTriggerChecks(_param0 []string) ([]*moira_alert.TriggerCheck, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerChecks", _param0)
	ret0, _ := ret[0].([]*moira_alert.TriggerCheck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerChecks indicates an expected call of GetTriggerChecks
func (_mr *MockDatabaseMockRecorder) GetTriggerChecks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerChecks", arg0)
}

// GetTriggerIDs mocks base method
func (_m *MockDatabase) GetTriggerIDs() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerIDs indicates an expected call of GetTriggerIDs
func (_mr *MockDatabaseMockRecorder) GetTriggerIDs() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerIDs")
}

// GetTriggerLastCheck mocks base method
func (_m *MockDatabase) GetTriggerLastCheck(_param0 string) (moira_alert.CheckData, error) {
	ret := _m.ctrl.Call(_m, "GetTriggerLastCheck", _param0)
	ret0, _ := ret[0].(moira_alert.CheckData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggerLastCheck indicates an expected call of GetTriggerLastCheck
func (_mr *MockDatabaseMockRecorder) GetTriggerLastCheck(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerLastCheck", arg0)
}

// GetTriggerThrottling mocks base method
func (_m *MockDatabase) GetTriggerThrottling(_param0 string) (time.Time, time.Time) {
	ret := _m.ctrl.Call(_m, "GetTriggerThrottling", _param0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(time.Time)
	return ret0, ret1
}

// GetTriggerThrottling indicates an expected call of GetTriggerThrottling
func (_mr *MockDatabaseMockRecorder) GetTriggerThrottling(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggerThrottling", arg0)
}

// GetTriggers mocks base method
func (_m *MockDatabase) GetTriggers(_param0 []string) ([]*moira_alert.Trigger, error) {
	ret := _m.ctrl.Call(_m, "GetTriggers", _param0)
	ret0, _ := ret[0].([]*moira_alert.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriggers indicates an expected call of GetTriggers
func (_mr *MockDatabaseMockRecorder) GetTriggers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTriggers", arg0)
}

// GetUserContactIDs mocks base method
func (_m *MockDatabase) GetUserContactIDs(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetUserContactIDs", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserContactIDs indicates an expected call of GetUserContactIDs
func (_mr *MockDatabaseMockRecorder) GetUserContactIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserContactIDs", arg0)
}

// GetUserSubscriptionIDs mocks base method
func (_m *MockDatabase) GetUserSubscriptionIDs(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetUserSubscriptionIDs", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscriptionIDs indicates an expected call of GetUserSubscriptionIDs
func (_mr *MockDatabaseMockRecorder) GetUserSubscriptionIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserSubscriptionIDs", arg0)
}

// PushNotificationEvent mocks base method
func (_m *MockDatabase) PushNotificationEvent(_param0 *moira_alert.NotificationEvent, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "PushNotificationEvent", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushNotificationEvent indicates an expected call of PushNotificationEvent
func (_mr *MockDatabaseMockRecorder) PushNotificationEvent(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PushNotificationEvent", arg0, arg1)
}

// RegisterBotIfAlreadyNot mocks base method
func (_m *MockDatabase) RegisterBotIfAlreadyNot(_param0 string) bool {
	ret := _m.ctrl.Call(_m, "RegisterBotIfAlreadyNot", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RegisterBotIfAlreadyNot indicates an expected call of RegisterBotIfAlreadyNot
func (_mr *MockDatabaseMockRecorder) RegisterBotIfAlreadyNot(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterBotIfAlreadyNot", arg0)
}

// RemoveContact mocks base method
func (_m *MockDatabase) RemoveContact(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveContact", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContact indicates an expected call of RemoveContact
func (_mr *MockDatabaseMockRecorder) RemoveContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContact", arg0)
}

// RemoveMetricValues mocks base method
func (_m *MockDatabase) RemoveMetricValues(_param0 string, _param1 int64) error {
	ret := _m.ctrl.Call(_m, "RemoveMetricValues", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMetricValues indicates an expected call of RemoveMetricValues
func (_mr *MockDatabaseMockRecorder) RemoveMetricValues(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveMetricValues", arg0, arg1)
}

// RemoveNotification mocks base method
func (_m *MockDatabase) RemoveNotification(_param0 string) (int64, error) {
	ret := _m.ctrl.Call(_m, "RemoveNotification", _param0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveNotification indicates an expected call of RemoveNotification
func (_mr *MockDatabaseMockRecorder) RemoveNotification(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveNotification", arg0)
}

// RemovePattern mocks base method
func (_m *MockDatabase) RemovePattern(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemovePattern", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePattern indicates an expected call of RemovePattern
func (_mr *MockDatabaseMockRecorder) RemovePattern(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePattern", arg0)
}

// RemovePatternTriggerIDs mocks base method
func (_m *MockDatabase) RemovePatternTriggerIDs(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemovePatternTriggerIDs", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePatternTriggerIDs indicates an expected call of RemovePatternTriggerIDs
func (_mr *MockDatabaseMockRecorder) RemovePatternTriggerIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePatternTriggerIDs", arg0)
}

// RemovePatternWithMetrics mocks base method
func (_m *MockDatabase) RemovePatternWithMetrics(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemovePatternWithMetrics", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePatternWithMetrics indicates an expected call of RemovePatternWithMetrics
func (_mr *MockDatabaseMockRecorder) RemovePatternWithMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePatternWithMetrics", arg0)
}

// RemovePatternsMetrics mocks base method
func (_m *MockDatabase) RemovePatternsMetrics(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "RemovePatternsMetrics", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePatternsMetrics indicates an expected call of RemovePatternsMetrics
func (_mr *MockDatabaseMockRecorder) RemovePatternsMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePatternsMetrics", arg0)
}

// RemoveSubscription mocks base method
func (_m *MockDatabase) RemoveSubscription(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveSubscription", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSubscription indicates an expected call of RemoveSubscription
func (_mr *MockDatabaseMockRecorder) RemoveSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveSubscription", arg0)
}

// RemoveTag mocks base method
func (_m *MockDatabase) RemoveTag(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveTag", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTag indicates an expected call of RemoveTag
func (_mr *MockDatabaseMockRecorder) RemoveTag(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveTag", arg0)
}

// RemoveTrigger mocks base method
func (_m *MockDatabase) RemoveTrigger(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveTrigger", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTrigger indicates an expected call of RemoveTrigger
func (_mr *MockDatabaseMockRecorder) RemoveTrigger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveTrigger", arg0)
}

// SaveContact mocks base method
func (_m *MockDatabase) SaveContact(_param0 *moira_alert.ContactData) error {
	ret := _m.ctrl.Call(_m, "SaveContact", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveContact indicates an expected call of SaveContact
func (_mr *MockDatabaseMockRecorder) SaveContact(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveContact", arg0)
}

// SaveMetrics mocks base method
func (_m *MockDatabase) SaveMetrics(_param0 map[string]*moira_alert.MatchedMetric) error {
	ret := _m.ctrl.Call(_m, "SaveMetrics", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveMetrics indicates an expected call of SaveMetrics
func (_mr *MockDatabaseMockRecorder) SaveMetrics(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveMetrics", arg0)
}

// SaveSubscription mocks base method
func (_m *MockDatabase) SaveSubscription(_param0 *moira_alert.SubscriptionData) error {
	ret := _m.ctrl.Call(_m, "SaveSubscription", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSubscription indicates an expected call of SaveSubscription
func (_mr *MockDatabaseMockRecorder) SaveSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveSubscription", arg0)
}

// SaveSubscriptions mocks base method
func (_m *MockDatabase) SaveSubscriptions(_param0 []*moira_alert.SubscriptionData) error {
	ret := _m.ctrl.Call(_m, "SaveSubscriptions", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSubscriptions indicates an expected call of SaveSubscriptions
func (_mr *MockDatabaseMockRecorder) SaveSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveSubscriptions", arg0)
}

// SaveTrigger mocks base method
func (_m *MockDatabase) SaveTrigger(_param0 string, _param1 *moira_alert.Trigger) error {
	ret := _m.ctrl.Call(_m, "SaveTrigger", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTrigger indicates an expected call of SaveTrigger
func (_mr *MockDatabaseMockRecorder) SaveTrigger(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveTrigger", arg0, arg1)
}

// SetTriggerCheckLock mocks base method
func (_m *MockDatabase) SetTriggerCheckLock(_param0 string) (bool, error) {
	ret := _m.ctrl.Call(_m, "SetTriggerCheckLock", _param0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTriggerCheckLock indicates an expected call of SetTriggerCheckLock
func (_mr *MockDatabaseMockRecorder) SetTriggerCheckLock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerCheckLock", arg0)
}

// SetTriggerCheckMetricsMaintenance mocks base method
func (_m *MockDatabase) SetTriggerCheckMetricsMaintenance(_param0 string, _param1 map[string]int64) error {
	ret := _m.ctrl.Call(_m, "SetTriggerCheckMetricsMaintenance", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerCheckMetricsMaintenance indicates an expected call of SetTriggerCheckMetricsMaintenance
func (_mr *MockDatabaseMockRecorder) SetTriggerCheckMetricsMaintenance(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerCheckMetricsMaintenance", arg0, arg1)
}

// SetTriggerLastCheck mocks base method
func (_m *MockDatabase) SetTriggerLastCheck(_param0 string, _param1 *moira_alert.CheckData) error {
	ret := _m.ctrl.Call(_m, "SetTriggerLastCheck", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerLastCheck indicates an expected call of SetTriggerLastCheck
func (_mr *MockDatabaseMockRecorder) SetTriggerLastCheck(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerLastCheck", arg0, arg1)
}

// SetTriggerThrottling mocks base method
func (_m *MockDatabase) SetTriggerThrottling(_param0 string, _param1 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetTriggerThrottling", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTriggerThrottling indicates an expected call of SetTriggerThrottling
func (_mr *MockDatabaseMockRecorder) SetTriggerThrottling(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTriggerThrottling", arg0, arg1)
}

// SetUsernameID mocks base method
func (_m *MockDatabase) SetUsernameID(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "SetUsernameID", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUsernameID indicates an expected call of SetUsernameID
func (_mr *MockDatabaseMockRecorder) SetUsernameID(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetUsernameID", arg0, arg1, arg2)
}

// SubscribeMetricEvents mocks base method
func (_m *MockDatabase) SubscribeMetricEvents(_param0 *tomb_v2.Tomb) (<-chan *moira_alert.MetricEvent, error) {
	ret := _m.ctrl.Call(_m, "SubscribeMetricEvents", _param0)
	ret0, _ := ret[0].(<-chan *moira_alert.MetricEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeMetricEvents indicates an expected call of SubscribeMetricEvents
func (_mr *MockDatabaseMockRecorder) SubscribeMetricEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubscribeMetricEvents", arg0)
}

// UpdateMetricsHeartbeat mocks base method
func (_m *MockDatabase) UpdateMetricsHeartbeat() error {
	ret := _m.ctrl.Call(_m, "UpdateMetricsHeartbeat")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetricsHeartbeat indicates an expected call of UpdateMetricsHeartbeat
func (_mr *MockDatabaseMockRecorder) UpdateMetricsHeartbeat() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMetricsHeartbeat")
}
