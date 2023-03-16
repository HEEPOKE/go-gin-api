package auth

import "Backend/go-api/model"

type UserService interface {
	CheckUserExistence(username string) (bool, error)
	CreateUser(user *model.User) error
}
