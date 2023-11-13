package services

import (
	"dineflow-api-gateway/model"
)

type AuthService interface {
	SignUpUser(*model.SignUpInput) (*model.DBResponse, error)
	SignInUser(*model.SignInInput) (*model.DBResponse, error)
}
