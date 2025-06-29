package services

import "7solutions/backend/core/models"

type UserService interface {
	CreateUser(payload models.SrvCreateUserModel) (result models.Response)

	GetUserByID(id string) (result models.Response)

	SignIn(payload models.SrvSignInModel) (result models.Response)

	Gets() (result models.Response)

	UpdateUser(id string, payload models.SrvUpdateUserModel) (result models.Response)

	DeleteUser(id string) (result models.Response)

	CountUser() (result models.Response)
}
