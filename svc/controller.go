package svc

import (
	"bitgo/entity"
	"fmt"
)

type Controller struct {
	*NotificationSvc
	*UserSvc
}

func (c *Controller) CreateNotification(userId int, content string) (int, error) {
	notification, err := c.NotificationSvc.CreateNotification(userId, content)
	if err != nil {
		return 0, err
	}
	return notification.GetId(), nil
}

func (c *Controller) SentNotification(notificationId int) error {
	err := c.NotificationSvc.SendNotification(notificationId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) PrintNotifications(userId int, status entity.NotificationStatus) ([]*entity.Notification, error) {
	notfications, err := c.NotificationSvc.GetNotifications(userId, status)
	if err != nil {
		return nil, err
	}
	for _, n := range notfications {
		fmt.Println(n)
	}
	return notfications, nil
}
