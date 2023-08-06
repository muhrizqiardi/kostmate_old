package service

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/kostmate/common/pkg/entities"
	"github.com/muhrizqiardi/kostmate/userservice/pkg/db_query"
	"github.com/muhrizqiardi/kostmate/userservice/pkg/dtos"
)

type Service struct {
	dbq db_query.DBQuerier
}

type Servicer interface {
	CreateNewUser(payload dtos.CreateNewUserDTO) (entities.UserEntity, error)
	GetOneUserByID(id uuid.UUID) (entities.UserEntity, error)
	UpdateOneUserByID(id uuid.UUID, payload dtos.UpdateOneUserByIDDTO) (entities.UserEntity, error)
	DeleteOneUserByID(id uuid.UUID) (entities.UserEntity, error)
}

func NewService(dbq db_query.DBQuerier) *Service {
	return &Service{dbq}
}

func (s *Service) CreateNewUser(payload dtos.CreateNewUserDTO) (entities.UserEntity, error) {
	newUser, err := s.dbq.InsertNewUser(payload.Email, payload.FullName, payload.Role)
	if err != nil {
		return entities.UserEntity{}, err
	}

	return newUser, nil
}

func (s *Service) GetOneUserByID(id uuid.UUID) (entities.UserEntity, error) {
	user, err := s.dbq.GetOneUserByID(id)
	if err != nil {
		return entities.UserEntity{}, err
	}

	return user, nil
}

func (s *Service) UpdateOneUserByID(id uuid.UUID, payload dtos.UpdateOneUserByIDDTO) (entities.UserEntity, error) {
	updatedUser, err := s.dbq.UpdateOneUserByID(id, payload.Email, payload.FullName, payload.Role)
	if err != nil {
		return entities.UserEntity{}, err
	}

	return updatedUser, nil
}

func (s *Service) DeleteOneUserByID(id uuid.UUID) (entities.UserEntity, error) {
	user, err := s.dbq.DeleteOneUserByID(id)
	if err != nil {
		return entities.UserEntity{}, err
	}

	return user, nil
}
