package svc

import (
	"bitgo/entity"
	"errors"
)

type UserSvc struct {
	repo map[int]*entity.User
}

func NewUserSvc() *UserSvc {
	return &UserSvc{repo: map[int]*entity.User{}}
}

func (u *UserSvc) GetUser(userId int) (*entity.User, error) {
	if user, ok := u.repo[userId]; ok {
		return user, nil
	}
	return nil, errors.New("not found")
}

func (u *UserSvc) CreateUser(preferences []entity.Channel, phoneNumber string) (*entity.User, error) {
	user := entity.NewUser(preferences, phoneNumber)
	u.repo[user.GetId()] = user
	return user, nil
}

func (u *UserSvc) UpdateUserPreferences(userId int, channels []entity.Channel) error {
	user, err := u.GetUser(userId)
	if err != nil {
		return err
	}
	user.SetPreferredChannels(channels)
	return nil
}
