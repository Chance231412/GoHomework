package api

import "homework_08/app/api/user"

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}
