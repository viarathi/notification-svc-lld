package svc

import (
	"bitgo/entity"
	"bitgo/notification_executor"
	"errors"
)

type NotificationSvc struct {
	userSvc *UserSvc
	repo    map[int]*entity.Notification
}

func NewNotificationSvc(userSvc *UserSvc) *NotificationSvc {
	return &NotificationSvc{userSvc: userSvc, repo: map[int]*entity.Notification{}}
}

func (n *NotificationSvc) CreateNotification(userId int, content string) (*entity.Notification, error) {
	notification := entity.NewNotification(userId, content)
	n.repo[notification.GetId()] = notification
	return notification, nil
}

func (n *NotificationSvc) getNotification(notificationId int) (*entity.Notification, error) {
	if notification, ok := n.repo[notificationId]; ok {
		return notification, nil
	}
	return nil, errors.New("notifications not found")
}

func (n *NotificationSvc) SendNotification(notificationId int) error {
	notification, err := n.getNotification(notificationId)
	if err != nil {
		return err
	}

	user, err := n.userSvc.GetUser(notification.GetUserId())
	if err != nil {
		return err
	}
	for _, channel := range user.GetPreferredChannels() {
		notifier, err := notification_executor.GetExecutor(channel)
		if err != nil {
			notification.UpdateStatus(entity.FAILURE)
			return err
		}
		notifier.Send(user, notification.GetBody())
	}
	notification.UpdateStatus(entity.SUCCESS)
	return nil
}

func (n *NotificationSvc) GetNotifications(userId int, status entity.NotificationStatus) ([]*entity.Notification, error) {
	var filteredNotifications []*entity.Notification
	for _, notification := range n.repo {
		if notification.GetUserId() == userId && status == notification.GetStatus() {
			filteredNotifications = append(filteredNotifications, notification)
		}
	}
	return filteredNotifications, nil

}
