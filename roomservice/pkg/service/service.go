package service

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/kostmate/common/pkg/entities"
	"github.com/muhrizqiardi/kostmate/roomservice/pkg/dbquery"
	"github.com/muhrizqiardi/kostmate/roomservice/pkg/dtos"
)

type Service struct {
	dbq dbquery.DBQuerier
}

type Servicer interface {
	CreateNewRoom(payload dtos.CreateNewRoomDTO) (entities.RoomEntity, error)
	GetOneRoomByID(id uuid.UUID) (entities.RoomEntity, error)
	GetOneRoomByUniqueName(uniqueName string) (entities.RoomEntity, error)
	GetManyRooms(payload dtos.GetManyRoomDTO) ([]entities.RoomEntity, error)
	UpdateOneRoomByID(id uuid.UUID, payload dtos.UpdateOneRoomByIDDTO) (entities.RoomEntity, error)
	DeleteOneRoomByID(id uuid.UUID) (entities.RoomEntity, error)
}

func NewService(dbq dbquery.DBQuerier) *Service {
	return &Service{dbq}
}

func (s *Service) CreateNewRoom(payload dtos.CreateNewRoomDTO) (entities.RoomEntity, error) {
	newRoom, err := s.dbq.InsertNewRoom(payload.UniqueName)
	if err != nil {
		return entities.RoomEntity{}, err
	}

	return newRoom, nil
}

func (s *Service) GetOneRoomByID(id uuid.UUID) (entities.RoomEntity, error) {
	room, err := s.dbq.GetOneRoomByID(id)
	if err != nil {
		return entities.RoomEntity{}, nil
	}

	return room, nil
}

func (s *Service) GetOneRoomByUniqueName(uniqueName string) (entities.RoomEntity, error) {
	room, err := s.dbq.GetOneRoomByUniqueName(uniqueName)
	if err != nil {
		return entities.RoomEntity{}, nil
	}

	return room, nil
}

func (s *Service) GetManyRooms(payload dtos.GetManyRoomDTO) ([]entities.RoomEntity, error) {
	rooms, err := s.dbq.GetManyRooms(payload.Limit, payload.Offset)
	if err != nil {
		return []entities.RoomEntity{}, nil
	}

	return rooms, nil
}

func (s *Service) UpdateOneRoomByID(id uuid.UUID, payload dtos.UpdateOneRoomByIDDTO) (entities.RoomEntity, error) {
	updatedRoom, err := s.dbq.UpdateOneRoomByID(id, payload.UniqueName)
	if err != nil {
		return entities.RoomEntity{}, nil
	}

	return updatedRoom, nil
}

func (s *Service) DeleteOneRoomByID(id uuid.UUID) (entities.RoomEntity, error) {
	deletedRoom, err := s.dbq.DeleteOneRoomByID(id)
	if err != nil {
		return entities.RoomEntity{}, nil
	}

	return deletedRoom, nil
}
