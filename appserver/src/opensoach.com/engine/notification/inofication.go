package notification

import (
	gnotmodels "opensoach.com/models/notification"
)

type INotificationEmail interface {
	GetEmailOptions() gnotmodels.EmailOptions
	GetNotificationConfig() gnotmodels.NotificationConfigModel
}
