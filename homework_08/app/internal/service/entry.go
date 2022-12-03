package service

import "homework_08/app/internal/service/user"

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}
