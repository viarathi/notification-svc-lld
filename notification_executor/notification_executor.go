package notification_executor

import (
	"bitgo/entity"
	"errors"
)

type NotificationExecutor interface {
	Send(user *entity.User, content string) error
}

func GetExecutor(channel entity.Channel) (NotificationExecutor, error) {
	if channel == entity.SMS {
		return &SmsNotificationExecutor{}, nil
	} else if channel == entity.EMAIL {
		return &EmailNotificationExecutor{}, nil
	}
	return nil, errors.New("Implemetor not found")
}
