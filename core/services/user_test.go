package services_test

import (
	"7solutions/backend/common/authorization"
	"7solutions/backend/core/models"
	"7solutions/backend/core/repositories"
	"7solutions/backend/core/services"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateUser(t *testing.T) {
	type test struct {
		Name  string
		Input models.SrvCreateUserModel
		Mock  struct {
			CreateUser struct {
				Input  models.RepoCreateUserModel
				Output models.RepoResUserModel
				Error  error
			}
		}
		Output models.Response
	}
	id := uuid.New().String()
	date := time.Now()
	cases := []test{
		{
			Name: "create user success",
			Input: models.SrvCreateUserModel{
				Name:     "bank",
				Email:    "test@test.com",
				Password: "123456",
			},
			Mock: struct {
				CreateUser struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}
			}{
				CreateUser: struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}{
					Input: models.RepoCreateUserModel{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "123456",
						CreateAt: date,
					},
					Output: models.RepoResUserModel{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "123456",
						CreateAt: date,
					},
					Error: nil,
				},
			},
			Output: models.Response{
				Status:  true,
				Message: "create user success",
				Code:    201,
				Data: models.SrvResUserModel{
					ID:       id,
					Name:     "bank",
					Email:    "test@test.com",
					Password: "123456",
					CreateAt: date.Format("2006-01-02 15:04:05"),
				},
			},
		},
		{
			Name: "error name not found",
			Input: models.SrvCreateUserModel{
				Name:     "",
				Email:    "test@test.com",
				Password: "123456",
			},
			Mock: struct {
				CreateUser struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}
			}{
				CreateUser: struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "name is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error email not found",
			Input: models.SrvCreateUserModel{
				Name:     "bank",
				Email:    "",
				Password: "123456",
			},
			Mock: struct {
				CreateUser struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}
			}{
				CreateUser: struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "email is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error password not found",
			Input: models.SrvCreateUserModel{
				Name:     "bank",
				Email:    "test@test.com",
				Password: "",
			},
			Mock: struct {
				CreateUser struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}
			}{
				CreateUser: struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "password is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error email invalid",
			Input: models.SrvCreateUserModel{
				Name:     "bank",
				Email:    "test",
				Password: "123456",
			},
			Mock: struct {
				CreateUser struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}
			}{
				CreateUser: struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "email invalid",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error create user",
			Input: models.SrvCreateUserModel{
				Name:     "bank",
				Email:    "test@test.com",
				Password: "123456",
			},
			Mock: struct {
				CreateUser struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}
			}{
				CreateUser: struct {
					Input  models.RepoCreateUserModel
					Output models.RepoResUserModel
					Error  error
				}{
					Input: models.RepoCreateUserModel{
						Name:     "bank",
						Email:    "test@test.com",
						Password: "123456",
					},
					Output: models.RepoResUserModel{},
					Error:  errors.New("error create user"),
				},
			},
			Output: models.Response{
				Status:  false,
				Message: "error create user",
				Code:    400,
				Data:    nil,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			auth := authorization.NewAuthorizationMock()
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("CreateUser", mock.AnythingOfType("models.RepoCreateUserModel")).Return(c.Mock.CreateUser.Output, c.Mock.CreateUser.Error)
			userSrv := services.NewUserService(auth, userRepo)

			result := userSrv.CreateUser(c.Input)
			assert.Equal(t, result, c.Output)
		})
	}
}

func Test_GetUserByID(t *testing.T) {
	type test struct {
		Name  string
		Input string
		Mock  struct {
			GetUserByID struct {
				Input  string
				Output models.RepoResUserModel
				Error  error
			}
		}
		Output models.Response
	}
	id := uuid.New().String()
	date := time.Now()
	cases := []test{
		{
			Name:  "get user success",
			Input: id,
			Mock: struct {
				GetUserByID struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
			}{
				GetUserByID: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{
					Input: id,
					Output: models.RepoResUserModel{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "123456",
						CreateAt: date,
					},
					Error: nil,
				},
			},
			Output: models.Response{
				Status:  true,
				Message: "get user success",
				Code:    200,
				Data: models.SrvResUserModel{
					ID:       id,
					Name:     "bank",
					Email:    "test@test.com",
					Password: "123456",
					CreateAt: date.Format("2006-01-02 15:04:05"),
				},
			},
		},
		{
			Name:  "error id not found",
			Input: "",
			Mock: struct {
				GetUserByID struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
			}{
				GetUserByID: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "id is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name:  "error get user",
			Input: id,
			Mock: struct {
				GetUserByID struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
			}{
				GetUserByID: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{
					Input:  id,
					Output: models.RepoResUserModel{},
					Error:  errors.New("error get user"),
				},
			},
			Output: models.Response{
				Status:  false,
				Message: "error get user",
				Code:    400,
				Data:    nil,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			auth := authorization.NewAuthorizationMock()
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("GetUserByID", c.Mock.GetUserByID.Input).Return(c.Mock.GetUserByID.Output, c.Mock.GetUserByID.Error)
			userSrv := services.NewUserService(auth, userRepo)

			result := userSrv.GetUserByID(c.Input)
			assert.Equal(t, result, c.Output)
		})
	}
}

func Test_SignIn(t *testing.T) {
	type test struct {
		Name  string
		Input models.SrvSignInModel
		Mock  struct {
			GetUserByEmail struct {
				Input  string
				Output models.RepoResUserModel
				Error  error
			}
			GenerateToken struct {
				Input  authorization.AppAuthorizationClaim
				Output string
				Error  error
			}
		}
		Output models.Response
	}
	id := uuid.New().String()
	cases := []test{
		{
			Name: "sign in success",
			Input: models.SrvSignInModel{
				Email:    "test@test.com",
				Password: "123456",
			},
			Mock: struct {
				GetUserByEmail struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
				GenerateToken struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}
			}{
				GetUserByEmail: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{
					Input: "test@test.com",
					Output: models.RepoResUserModel{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "$2a$10$eZjqtJ6RE6ALsVLLo6cfz.JYIkwLTlB3HV1xAvk4Im2d98uvuUKMq",
					},
					Error: nil,
				},
				GenerateToken: struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}{
					Input: authorization.AppAuthorizationClaim{
						UserId:   id,
						Audience: "7solutions",
						Issuer:   "7solutions",
					},
					Output: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI3c29sdXRpb25zIiwiZXhwIjoxNzUxMTAxMDEzLCJpYXQiOjE3NTEwMTQ2MTMsImlzcyI6Ijdzb2x1dGlvbnMiLCJzdWIiOiIyMjdiNmVkZS05NWYxLTRhMDEtOTZkZi01YzRmZmI2MTA2M2MifQ.jEM1-iBUL4j0Jt5Q_Xpi3MqmhZ_3loXGJywiyr_eKjE",
					Error:  nil,
				},
			},
			Output: models.Response{
				Status:  true,
				Message: "sign in success",
				Code:    200,
				Data: models.SrvSignInResModel{
					Type:        "Bearer",
					AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI3c29sdXRpb25zIiwiZXhwIjoxNzUxMTAxMDEzLCJpYXQiOjE3NTEwMTQ2MTMsImlzcyI6Ijdzb2x1dGlvbnMiLCJzdWIiOiIyMjdiNmVkZS05NWYxLTRhMDEtOTZkZi01YzRmZmI2MTA2M2MifQ.jEM1-iBUL4j0Jt5Q_Xpi3MqmhZ_3loXGJywiyr_eKjE",
				},
			},
		},
		{
			Name: "error email not found",
			Input: models.SrvSignInModel{
				Email:    "",
				Password: "123456",
			},
			Mock: struct {
				GetUserByEmail struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
				GenerateToken struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}
			}{
				GetUserByEmail: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{},
				GenerateToken: struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "email is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error password not found",
			Input: models.SrvSignInModel{
				Email:    "test@test.com",
				Password: "",
			},
			Mock: struct {
				GetUserByEmail struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
				GenerateToken struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}
			}{
				GetUserByEmail: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{},
				GenerateToken: struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "password is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error email not valid",
			Input: models.SrvSignInModel{
				Email:    "test",
				Password: "123456",
			},
			Mock: struct {
				GetUserByEmail struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
				GenerateToken struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}
			}{
				GetUserByEmail: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{},
				GenerateToken: struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "email invalid",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error get user by email",
			Input: models.SrvSignInModel{
				Email:    "test@test.com",
				Password: "123456",
			},
			Mock: struct {
				GetUserByEmail struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
				GenerateToken struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}
			}{
				GetUserByEmail: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{
					Input:  "test@test.com",
					Output: models.RepoResUserModel{},
					Error:  errors.New("error get user by email"),
				},
				GenerateToken: struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "error get user by email",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error password not match",
			Input: models.SrvSignInModel{
				Email:    "test@test.com",
				Password: "123456",
			},
			Mock: struct {
				GetUserByEmail struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
				GenerateToken struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}
			}{
				GetUserByEmail: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{
					Input: "test@test.com",
					Output: models.RepoResUserModel{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "$2a$10$eZjqtJ6RE6ALsVLLo6cfz.JYIkwLTlB3HV1xAvk4Im2d98uvu",
					},
					Error: nil,
				},
				GenerateToken: struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "invalid password",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error generate token",
			Input: models.SrvSignInModel{
				Email:    "test@test.com",
				Password: "123456",
			},
			Mock: struct {
				GetUserByEmail struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}
				GenerateToken struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}
			}{
				GetUserByEmail: struct {
					Input  string
					Output models.RepoResUserModel
					Error  error
				}{
					Input: "test@test.com",
					Output: models.RepoResUserModel{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "$2a$10$eZjqtJ6RE6ALsVLLo6cfz.JYIkwLTlB3HV1xAvk4Im2d98uvuUKMq",
					},
					Error: nil,
				},
				GenerateToken: struct {
					Input  authorization.AppAuthorizationClaim
					Output string
					Error  error
				}{
					Input: authorization.AppAuthorizationClaim{
						UserId:   id,
						Audience: "7solutions",
						Issuer:   "7solutions",
					},
					Output: "",
					Error:  errors.New("error generate token"),
				},
			},
			Output: models.Response{
				Status:  false,
				Message: "error generate token",
				Code:    400,
				Data:    nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			auth := authorization.NewAuthorizationMock()
			auth.On("GenerateToken", c.Mock.GenerateToken.Input).Return(c.Mock.GenerateToken.Output, c.Mock.GenerateToken.Error)
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("GetUserByEmail", c.Mock.GetUserByEmail.Input).Return(c.Mock.GetUserByEmail.Output, c.Mock.GetUserByEmail.Error)
			userSrv := services.NewUserService(auth, userRepo)

			result := userSrv.SignIn(c.Input)
			assert.Equal(t, result, c.Output)
		})
	}
}

func Test_Gets(t *testing.T) {
	type test struct {
		Name  string
		Input string
		Mock  struct {
			GetUsers struct {
				Output []models.RepoResUserModel
				Error  error
			}
		}
		Output models.Response
	}
	id := uuid.New().String()
	cases := []test{
		{
			Name: "gets success",
			Mock: struct {
				GetUsers struct {
					Output []models.RepoResUserModel
					Error  error
				}
			}{
				GetUsers: struct {
					Output []models.RepoResUserModel
					Error  error
				}{
					Output: []models.RepoResUserModel{
						{
							ID:       id,
							Name:     "bank",
							Email:    "test@test.com",
							Password: "$2a$10$eZjqtJ6RE6ALsVLLo6cfz.JYIkwLTlB3HV1xAvk4Im2d98uvuUKMq",
						},
					},
					Error: nil,
				},
			},
			Output: models.Response{
				Status:  true,
				Message: "get users success",
				Code:    200,
				Data: []models.RepoResUserModel{
					{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "$2a$10$eZjqtJ6RE6ALsVLLo6cfz.JYIkwLTlB3HV1xAvk4Im2d98uvuUKMq",
					},
				},
			},
		},
		{
			Name: "error get users",
			Mock: struct {
				GetUsers struct {
					Output []models.RepoResUserModel
					Error  error
				}
			}{
				GetUsers: struct {
					Output []models.RepoResUserModel
					Error  error
				}{
					Output: nil,
					Error:  errors.New("error get users"),
				},
			},
			Output: models.Response{
				Status:  false,
				Message: "error get users",
				Code:    400,
				Data:    nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			auth := authorization.NewAuthorizationMock()
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("GetUsers").Return(c.Mock.GetUsers.Output, c.Mock.GetUsers.Error)
			userSrv := services.NewUserService(auth, userRepo)

			result := userSrv.Gets()
			assert.Equal(t, result, c.Output)
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	type test struct {
		Name  string
		Input struct {
			ID      string
			Payload models.SrvUpdateUserModel
		}
		Mock struct {
			UpdateUser struct {
				Input struct {
					ID      string
					Payload models.RepoUpdateUserModel
				}
				Output models.RepoResUserModel
				Error  error
			}
		}
		Output models.Response
	}
	id := uuid.New().String()
	cases := []test{
		{
			Name: "update user success",
			Input: struct {
				ID      string
				Payload models.SrvUpdateUserModel
			}{
				ID: id,
				Payload: models.SrvUpdateUserModel{
					Name:  "bank",
					Email: "test@test.com",
				},
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}{
					Input: struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}{
						ID: id,
						Payload: models.RepoUpdateUserModel{
							Name:  "bank",
							Email: "test@test.com",
						},
					},
					Output: models.RepoResUserModel{
						ID:       id,
						Name:     "bank",
						Email:    "test@test.com",
						Password: "$2a$10$eZjqtJ6RE6ALsVLLo6cfz.JYIkwLTlB3HV1xAvk4Im2d98uvuUKMq",
					},
					Error: nil,
				},
			},
			Output: models.Response{
				Status:  true,
				Message: "update user success",
				Code:    200,
				Data: models.RepoResUserModel{
					ID:       id,
					Name:     "bank",
					Email:    "test@test.com",
					Password: "$2a$10$eZjqtJ6RE6ALsVLLo6cfz.JYIkwLTlB3HV1xAvk4Im2d98uvuUKMq",
				},
			},
		},
		{
			Name: "error id not found",
			Input: struct {
				ID      string
				Payload models.SrvUpdateUserModel
			}{
				ID: "",
				Payload: models.SrvUpdateUserModel{
					Name:  "bank",
					Email: "test@test.com",
				},
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "id is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "email not valid",
			Input: struct {
				ID      string
				Payload models.SrvUpdateUserModel
			}{
				ID: id,
				Payload: models.SrvUpdateUserModel{
					Name:  "bank",
					Email: "test",
				},
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "email invalid",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name: "error update user",
			Input: struct {
				ID      string
				Payload models.SrvUpdateUserModel
			}{
				ID: id,
				Payload: models.SrvUpdateUserModel{
					Name:  "bank",
					Email: "test@test.com",
				},
			},
			Mock: struct {
				UpdateUser struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}
			}{
				UpdateUser: struct {
					Input struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}
					Output models.RepoResUserModel
					Error  error
				}{
					Input: struct {
						ID      string
						Payload models.RepoUpdateUserModel
					}{
						ID: id,
						Payload: models.RepoUpdateUserModel{
							Name:  "bank",
							Email: "test@test.com",
						},
					},
					Output: models.RepoResUserModel{},
					Error:  errors.New("error update user"),
				},
			},
			Output: models.Response{
				Status:  false,
				Message: "error update user",
				Code:    400,
				Data:    nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			auth := authorization.NewAuthorizationMock()
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("UpdateUser", c.Mock.UpdateUser.Input.ID, c.Mock.UpdateUser.Input.Payload).Return(c.Mock.UpdateUser.Output, c.Mock.UpdateUser.Error)
			userSrv := services.NewUserService(auth, userRepo)

			result := userSrv.UpdateUser(c.Input.ID, c.Input.Payload)
			assert.Equal(t, result, c.Output)
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	type test struct {
		Name  string
		Input string
		Mock  struct {
			DeleteUser struct {
				Input string
				Error error
			}
		}
		Output models.Response
	}
	id := uuid.New().String()
	cases := []test{
		{
			Name:  "delete user success",
			Input: id,
			Mock: struct {
				DeleteUser struct {
					Input string
					Error error
				}
			}{
				DeleteUser: struct {
					Input string
					Error error
				}{
					Input: id,
					Error: nil,
				},
			},
			Output: models.Response{
				Status:  true,
				Message: "delete user success",
				Code:    200,
				Data:    nil,
			},
		},
		{
			Name:  "error id not found",
			Input: "",
			Mock: struct {
				DeleteUser struct {
					Input string
					Error error
				}
			}{
				DeleteUser: struct {
					Input string
					Error error
				}{},
			},
			Output: models.Response{
				Status:  false,
				Message: "id is required",
				Code:    400,
				Data:    nil,
			},
		},
		{
			Name:  "error delete user",
			Input: id,
			Mock: struct {
				DeleteUser struct {
					Input string
					Error error
				}
			}{
				DeleteUser: struct {
					Input string
					Error error
				}{
					Input: id,
					Error: errors.New("error delete user"),
				},
			},
			Output: models.Response{
				Status:  false,
				Message: "error delete user",
				Code:    400,
				Data:    nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			auth := authorization.NewAuthorizationMock()
			userRepo := repositories.NewUserRepositoryMock()
			userRepo.On("DeleteUser", c.Mock.DeleteUser.Input).Return(c.Mock.DeleteUser.Error)
			userSrv := services.NewUserService(auth, userRepo)

			result := userSrv.DeleteUser(c.Input)
			assert.Equal(t, result, c.Output)
		})
	}
}
