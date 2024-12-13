package entity

type User struct {
	id                int
	preferredChannels []Channel
	phoneNumber       string
}

var userId int = 0

func NewUser(preferredChannels []Channel, phone string) *User {
	// can create fns around this
	userId++
	return &User{
		id:                userId,
		preferredChannels: preferredChannels,
		phoneNumber:       phone,
	}
}

func (u *User) GetPreferredChannels() []Channel {
	return u.preferredChannels
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetPhoneNumber() string {
	return u.phoneNumber
}

func (u *User) SetPreferredChannels(channels []Channel) {
	u.preferredChannels = channels
}
