package notification_executor

import (
	"bitgo/entity"
	"fmt"
)

type SmsNotificationExecutor struct {
}

func (s *SmsNotificationExecutor) Send(user *entity.User, body string) error {
	fmt.Printf("got sms request %v, for user %v \n", body, user.GetPhoneNumber())
	return nil
}
