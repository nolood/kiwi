package userdto

import (
	"kiwi/.gen/kiwi/public/model"
)

type UserWithProfile struct {
	Profile model.Profiles
	User    model.Users
}

type ProfileUpdate struct {
	Age       *int
	Gender    *string
	About     *string
	PhotoId   *string
	Longitude *float64
	Latitude  *float64
}
