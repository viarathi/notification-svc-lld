package main

import (
	"bitgo/entity"
	"bitgo/svc"
)

/*
*
Functional Req
  - We should be able to create notification (accept channel, notification body, toWhom)
  - user preferences should be taken in account (ex: channels preferences, rate limits? (frqequency))
  - priority of events
  - History of notifications (user)
  - Update a notification status.

# APIs

POST /user
PATCH /user/:userId

	{
	  preferences : {}
	}

POST /notifications  : {userId}

	{
	  content: {}
	}

- CreateNotification in DB (created state) : notification_id

- Send(notification_details)
  - fetch user_preference by user_id
  - count (default_rate_limit, user_id, channel) : from notification
  - channel : channelImpletor : Send()

GET /notifications?limit=10&user_id=?status=?

# Entities

Notifications
- notification_id
- user_id
- channel_type
- body
- status (INTIATED)

::

Users
- user_id
- user_preferences (frequency_day, channels []channel_type)

config # number so
Channel_type -> template
- channel_id
- name (SMS/..)
- template
- how to actually send this (3rd party)
*/
func main() {

	// create user with preferences

	userSvc := svc.NewUserSvc()
	notificationSvc := svc.NewNotificationSvc(userSvc)
	controller := svc.Controller{
		NotificationSvc: notificationSvc,
		UserSvc:         userSvc,
	}

	// via controller
	user, _ := userSvc.CreateUser([]entity.Channel{
		entity.SMS,
	}, "1234")

	notificationId, _ := controller.CreateNotification(user.GetId(), "BTC price is now 90K..")
	controller.PrintNotifications(user.GetId(), entity.CREATED)

	controller.SendNotification(notificationId)

	controller.PrintNotifications(user.GetId(), entity.SUCCESS)

	userSvc.UpdateUserPreferences(user.GetId(), []entity.Channel{entity.EMAIL})

	notificationId, _ = controller.CreateNotification(user.GetId(), "BTC price is now 100K..")
	controller.PrintNotifications(user.GetId(), entity.SUCCESS)

	controller.SendNotification(notificationId)

	controller.PrintNotifications(user.GetId(), entity.SUCCESS)

}
