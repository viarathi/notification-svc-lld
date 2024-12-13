package notification_executor

import (
	"bitgo/entity"
	"fmt"
)

type EmailNotificationExecutor struct {
}

func (s *EmailNotificationExecutor) Send(user *entity.User, body string) error {
	fmt.Printf("got email request %v, for user %v \n", body, user.GetPhoneNumber())
	return nil
}
