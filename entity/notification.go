package entity

import "fmt"

type Channel string
type NotificationStatus string

const (
	EMAIL Channel = "email"
	SMS   Channel = "sms"
)

const (
	CREATED NotificationStatus = "created"
	SUCCESS NotificationStatus = "success"
	FAILURE NotificationStatus = "failure"
)

type Notification struct {
	id     int
	userId int
	body   string
	status NotificationStatus
}

func (n *Notification) String() string {
	return fmt.Sprintf("id: %d  status: %v", n.id, n.status)
}
func (n *Notification) UpdateStatus(status NotificationStatus) {
	n.status = status
}

func (n *Notification) GetId() int {
	return n.id
}
func (n *Notification) GetUserId() int {
	return n.userId
}
func (n *Notification) GetBody() string {
	return n.body
}

func (n *Notification) GetStatus() NotificationStatus {
	return n.status
}

var notificationId int = 0

func NewNotification(userId int, body string) *Notification {
	notificationId++
	return &Notification{
		id:     notificationId,
		userId: userId,
		body:   body,
		status: CREATED,
	}
}
