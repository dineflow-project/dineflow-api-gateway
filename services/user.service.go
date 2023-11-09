package services

import "api-gateway/model"

type UserService interface {
	FindUserById(string) (*model.DBResponse, error)
	FindUserByEmail(string) (*model.DBResponse, error)
}
