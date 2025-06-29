package models

import "time"

type RepoResUserModel struct {
	ID       string    `json:"id" bson:"id"`
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email"`
	Password string    `json:"password" bson:"password"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}

type RepoCreateUserModel struct {
	ID       string    `json:"id" bson:"id"`
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email"`
	Password string    `json:"password" bson:"password"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}

type SrvCreateUserModel struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type SrvResUserModel struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	CreateAt string `json:"createAt" bson:"createAt"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type RepoFilterUserModel struct {
	ID    string `json:"id" bson:"id"`
	Email string `json:"email" bson:"email"`
}

type SrvSignInModel struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type SrvSignInResModel struct {
	Type        string `json:"type" bson:"type"`
	AccessToken string `json:"accessToken" bson:"accessToken"`
}

type RepoUpdateUserModel struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Email string `json:"email" bson:"email,omitempty"`
}

type SrvUpdateUserModel struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
