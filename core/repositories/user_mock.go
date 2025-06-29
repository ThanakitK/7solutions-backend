package repositories

import (
	"7solutions/backend/core/models"

	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepoMock {
	return &userRepoMock{}
}

func (m *userRepoMock) CreateUser(payload models.RepoCreateUserModel) (result models.RepoResUserModel, err error) {
	args := m.Called(payload)
	return args.Get(0).(models.RepoResUserModel), args.Error(1)
}

func (m *userRepoMock) GetUserByID(id string) (result models.RepoResUserModel, err error) {
	args := m.Called(id)
	return args.Get(0).(models.RepoResUserModel), args.Error(1)
}

func (m *userRepoMock) GetUserByEmail(email string) (result models.RepoResUserModel, err error) {
	args := m.Called(email)
	return args.Get(0).(models.RepoResUserModel), args.Error(1)
}

func (m *userRepoMock) GetUsers() (result []models.RepoResUserModel, err error) {
	args := m.Called()
	return args.Get(0).([]models.RepoResUserModel), args.Error(1)
}

func (m *userRepoMock) UpdateUser(id string, payload models.RepoUpdateUserModel) (result models.RepoResUserModel, err error) {
	args := m.Called(id, payload)
	return args.Get(0).(models.RepoResUserModel), args.Error(1)
}

func (m *userRepoMock) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *userRepoMock) CountUser() (result int64, err error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
