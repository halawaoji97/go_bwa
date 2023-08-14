package user

import "time"

type User struct {
	Id             int
	Name           string
	Email          string
	Password       string
	Occupation     string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
