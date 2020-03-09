package playstore

// DeveloperNotification is sent by a Pub/Sub topic.
// Detailed description is following.
// https://developer.android.com/google/play/billing/realtime_developer_notifications.html#json_specification
type DeveloperNotification struct {
	Version                  string                   `json:"version" bson:"version"`
	PackageName              string                   `json:"packageName" bson:"package_name"`
	EventTimeMillis          string                   `json:"eventTimeMillis" bson:"event_time_millis"`
	SubscriptionNotification SubscriptionNotification `json:"subscriptionNotification,omitempty" bson:"subscription_notification,omitempty"`
	TestNotification         SubscriptionNotification `json:"testNotification,omitempty" bson:"test_notification,omitempty"`
}

// SubscriptionNotification has subscription status as notificationType, toke and subscription id
// to confirm status by calling Google Android Publisher API.
type SubscriptionNotification struct {
	Version          string           `json:"version" bson:"version"`
	NotificationType NotificationType `json:"notificationType,omitempty" bson:"notification_type,omitempty"`
	PurchaseToken    string           `json:"purchaseToken,omitempty" bson:"purchase_token,omitempty"`
	SubscriptionID   string           `json:"subscriptionId,omitempty" bson:"subscription_id,omitempty"`
}

type NotificationType int

const (
	NotificationTypeRecovered NotificationType = iota + 1
	NotificationTypeRenewed
	NotificationTypeCanceled
	NotificationTypePurchased
	NotificationTypeAccountHold
	NotificationTypeGracePeriod
	NotificationTypeReactivated
)
