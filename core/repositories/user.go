package repositories

import "7solutions/backend/core/models"

type UserRepository interface {
	CreateUser(payload models.RepoCreateUserModel) (result models.RepoResUserModel, err error)

	GetUserByID(id string) (result models.RepoResUserModel, err error)

	GetUserByEmail(email string) (result models.RepoResUserModel, err error)

	GetUsers() (result []models.RepoResUserModel, err error)

	UpdateUser(id string, payload models.RepoUpdateUserModel) (result models.RepoResUserModel, err error)

	DeleteUser(id string) error

	CountUser() (result int64, err error)
}
