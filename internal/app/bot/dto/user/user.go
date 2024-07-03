package userdto

import (
	"kiwi/.gen/kiwi/public/model"
)

type UserWithProfile struct {
	Profile model.Profiles
	User    model.Users
}
