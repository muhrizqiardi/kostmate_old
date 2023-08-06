package service

import (
	"errors"

	"github.com/muhrizqiardi/kostmate/common/pkg/entities"
	"github.com/muhrizqiardi/kostmate/userservice/pkg/db_query"
	"github.com/muhrizqiardi/kostmate/userservice/pkg/dtos"
)

type Service struct {
	dbq db_query.DBQuerier
}

type Servicer interface {
	CreateNewUser() (entities.UserEntity, error)
	// GetOneUserByID() (entities.UserEntity, error)
	// UpdateOneUserByID() (entities.UserEntity, error)
	// DeleteOneUserByID() (entities.UserEntity, error)
}

func NewService(dbq db_query.DBQuerier) *Service {
	return &Service{dbq}
}

func (s *Service) CreateNewUser(payload dtos.CreateNewUserDTO) (entities.UserEntity, error) {
	newUser, err := s.dbq.InsertNewUser(payload.Email, payload.FullName, payload.Role)
	if err != nil {
		return entities.UserEntity{}, errors.New("Failed to create new user")
	}

	return newUser, nil
}
