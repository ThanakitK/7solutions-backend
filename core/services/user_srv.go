package services

import (
	"7solutions/backend/common/authorization"
	"7solutions/backend/core/models"
	"7solutions/backend/core/repositories"
	"7solutions/backend/utils"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type userSrv struct {
	auth     authorization.AppAuthorization
	userRepo repositories.UserRepository
}

func NewUserService(auth authorization.AppAuthorization, userRepo repositories.UserRepository) UserService {
	return &userSrv{
		auth:     auth,
		userRepo: userRepo,
	}
}

func (s *userSrv) CreateUser(payload models.SrvCreateUserModel) (result models.Response) {
	if payload.Name == "" {
		return models.Response{
			Status:  false,
			Message: "name is required",
			Code:    400,
			Data:    nil,
		}
	}
	if payload.Email == "" {
		return models.Response{
			Status:  false,
			Message: "email is required",
			Code:    400,
			Data:    nil,
		}
	}
	if payload.Password == "" {
		return models.Response{
			Status:  false,
			Message: "password is required",
			Code:    400,
			Data:    nil,
		}
	}

	_, err := mail.ParseAddress(payload.Email)
	if err != nil {
		return models.Response{
			Status:  false,
			Message: "email invalid",
			Code:    400,
			Data:    nil,
		}
	}
	hashPassword, _ := utils.Bcryp_Encryption(payload.Password)

	payloadCreate := models.RepoCreateUserModel{
		ID:       uuid.New().String(),
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashPassword,
		CreateAt: time.Now(),
	}
	res, err := s.userRepo.CreateUser(payloadCreate)
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}
	data := models.SrvResUserModel{
		ID:       res.ID,
		Name:     res.Name,
		Email:    res.Email,
		Password: res.Password,
		CreateAt: res.CreateAt.Format("2006-01-02 15:04:05"),
	}
	result = models.Response{
		Status:  true,
		Message: "create user success",
		Code:    201,
		Data:    data,
	}
	return result
}

func (s *userSrv) GetUserByID(id string) (result models.Response) {
	if id == "" {
		return models.Response{
			Status:  false,
			Message: "id is required",
			Code:    400,
			Data:    nil,
		}
	}
	res, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}

	data := models.SrvResUserModel{
		ID:       res.ID,
		Name:     res.Name,
		Email:    res.Email,
		Password: res.Password,
		CreateAt: res.CreateAt.Format("2006-01-02 15:04:05"),
	}
	result = models.Response{
		Status:  true,
		Message: "get user success",
		Code:    200,
		Data:    data,
	}
	return result
}

func (s *userSrv) SignIn(payload models.SrvSignInModel) (result models.Response) {
	if payload.Email == "" {
		return models.Response{
			Status:  false,
			Message: "email is required",
			Code:    400,
			Data:    nil,
		}
	}
	if payload.Password == "" {
		return models.Response{
			Status:  false,
			Message: "password is required",
			Code:    400,
			Data:    nil,
		}
	}
	_, err := mail.ParseAddress(payload.Email)
	if err != nil {
		return models.Response{
			Status:  false,
			Message: "email invalid",
			Code:    400,
			Data:    nil,
		}
	}
	user, err := s.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}
	if !utils.Bcryp_Compare(user.Password, payload.Password) {
		return models.Response{
			Status:  false,
			Message: "invalid password",
			Code:    400,
			Data:    nil,
		}
	}

	accessToken, err := s.auth.GenerateToken(authorization.AppAuthorizationClaim{
		UserId:   user.ID,
		Audience: "7solutions",
		Issuer:   "7solutions",
	})
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}
	data := models.SrvSignInResModel{
		Type:        "Bearer",
		AccessToken: accessToken,
	}

	result = models.Response{
		Status:  true,
		Message: "sign in success",
		Code:    200,
		Data:    data,
	}
	return result
}

func (s *userSrv) Gets() (result models.Response) {
	res, err := s.userRepo.GetUsers()
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}
	result = models.Response{
		Status:  true,
		Message: "get users success",
		Code:    200,
		Data:    res,
	}
	return result
}

func (s *userSrv) UpdateUser(id string, payload models.SrvUpdateUserModel) (result models.Response) {
	if id == "" {
		return models.Response{
			Status:  false,
			Message: "id is required",
			Code:    400,
			Data:    nil,
		}
	}
	if payload.Email != "" {
		_, err := mail.ParseAddress(payload.Email)
		if err != nil {
			return models.Response{
				Status:  false,
				Message: "email invalid",
				Code:    400,
				Data:    nil,
			}
		}
	}
	payloadUpdate := models.RepoUpdateUserModel(payload)
	res, err := s.userRepo.UpdateUser(id, payloadUpdate)
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}
	result = models.Response{
		Status:  true,
		Message: "update user success",
		Code:    200,
		Data:    res,
	}
	return result
}

func (s *userSrv) DeleteUser(id string) (result models.Response) {
	if id == "" {
		return models.Response{
			Status:  false,
			Message: "id is required",
			Code:    400,
			Data:    nil,
		}
	}
	err := s.userRepo.DeleteUser(id)
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}
	result = models.Response{
		Status:  true,
		Message: "delete user success",
		Code:    200,
		Data:    nil,
	}
	return result
}

func (s *userSrv) CountUser() (result models.Response) {
	res, err := s.userRepo.CountUser()
	if err != nil {
		return models.Response{
			Status:  false,
			Message: err.Error(),
			Code:    400,
			Data:    nil,
		}
	}
	result = models.Response{
		Status:  true,
		Message: "count user success",
		Code:    200,
		Data:    res,
	}
	return result
}
