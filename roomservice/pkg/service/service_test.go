package service

import (
	"errors"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/muhrizqiardi/kostmate/common/pkg/entities"
	"github.com/muhrizqiardi/kostmate/roomservice/pkg/dtos"
)

type mockDBQuery struct {
	rooms []entities.RoomEntity
}

func newMockDBQuery() *mockDBQuery {
	return &mockDBQuery{}
}

func (mdbq *mockDBQuery) InsertNewRoom(uniqueName string) (entities.RoomEntity, error) {
	for _, room := range mdbq.rooms {
		if room.UniqueName == uniqueName {
			return entities.RoomEntity{}, errors.New("Unique name already exists")
		}
	}

	newUUID, _ := uuid.NewRandom()
	newRoom := entities.RoomEntity{
		Id:         newUUID,
		UniqueName: uniqueName,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	mdbq.rooms = append(mdbq.rooms, newRoom)
	return newRoom, nil
}

func (mdbq *mockDBQuery) GetOneRoomByID(id uuid.UUID) (entities.RoomEntity, error) {
	for _, room := range mdbq.rooms {
		if room.Id.String() == id.String() {
			return room, nil
		}
	}

	return entities.RoomEntity{}, errors.New("Room not found")
}

func (mdbq *mockDBQuery) GetOneRoomByUniqueName(uniqueName string) (entities.RoomEntity, error) {
	for _, room := range mdbq.rooms {
		if room.UniqueName == uniqueName {
			return room, nil
		}
	}

	return entities.RoomEntity{}, errors.New("Room not found")
}

func (mdbq *mockDBQuery) GetManyRooms(limit int, offset int) ([]entities.RoomEntity, error) {
	if offset > len(mdbq.rooms) {
		return []entities.RoomEntity{}, errors.New("Offset more than limit")
	}
	if limit+offset > len(mdbq.rooms) {
		return mdbq.rooms[offset : len(mdbq.rooms)-1], nil
	}

	return mdbq.rooms[offset-1 : limit+1], nil
}

func (mdbq *mockDBQuery) UpdateOneRoomByID(id uuid.UUID, uniqueName string) (entities.RoomEntity, error) {
	for i, room := range mdbq.rooms {
		if room.Id.String() == id.String() {
			mdbq.rooms[i].UniqueName = uniqueName
			mdbq.rooms[i].UpdatedAt = time.Now()
			return mdbq.rooms[i], nil
		}
	}

	return entities.RoomEntity{}, errors.New("Room not found")
}

func (mdbq *mockDBQuery) DeleteOneRoomByID(id uuid.UUID) (entities.RoomEntity, error) {
	var deletedRoom entities.RoomEntity
	var foundIndex int

	for i, room := range mdbq.rooms {
		if room.Id.String() == id.String() {
			deletedRoom = mdbq.rooms[i]
			foundIndex = i
		}
	}

	mdbq.rooms = append(mdbq.rooms[:foundIndex], mdbq.rooms[foundIndex+1:]...)

	return deletedRoom, nil
}

func TestService(t *testing.T) {
	t.Run("should create new room", func(t *testing.T) {
		mdbq := newMockDBQuery()
		s := NewService(mdbq)

		mockDto := dtos.CreateNewRoomDTO{
			UniqueName: faker.Word(),
		}
		got, gotErr := s.CreateNewRoom(mockDto)
		if gotErr != nil {
			t.Errorf("Expect no error, got an error")
		}
		if got.UniqueName != mockDto.UniqueName {
			t.Errorf("Expected got.UniqueName to be %s, got %s", mockDto.UniqueName, got.UniqueName)
		}
	})

	t.Run("should return error when unique name already exists", func(t *testing.T) {
		mdbq := newMockDBQuery()
		s := NewService(mdbq)

		mockDto := dtos.CreateNewRoomDTO{
			UniqueName: faker.Word(),
		}
		s.CreateNewRoom(mockDto)
		_, gotErr := s.CreateNewRoom(mockDto)
		if gotErr == nil {
			t.Errorf("Expect error, got no error")
		}
	})

	t.Run("should get one room by ID", func(t *testing.T) {
		mdbq := newMockDBQuery()
		s := NewService(mdbq)

		mockDto := dtos.CreateNewRoomDTO{
			UniqueName: faker.Word(),
		}
		newRoom, _ := s.CreateNewRoom(mockDto)
		got, gotErr := s.GetOneRoomByID(newRoom.Id)
		if gotErr != nil {
			t.Errorf("Expect no error, got an error")
		}
		if got.Id.String() != newRoom.Id.String() {
			t.Errorf("Expected got.UniqueName to be %s, got %s", newRoom.Id.String(), got.Id.String())
		}
	})

	t.Run("should get one room by its unique name", func(t *testing.T) {
		mdbq := newMockDBQuery()
		s := NewService(mdbq)

		mockDto := dtos.CreateNewRoomDTO{
			UniqueName: faker.Word(),
		}
		newRoom, _ := s.CreateNewRoom(mockDto)
		got, gotErr := s.GetOneRoomByUniqueName(newRoom.UniqueName)
		if gotErr != nil {
			t.Errorf("Expect no error, got an error")
		}
		if got.Id.String() != newRoom.Id.String() {
			t.Errorf("Expected got.UniqueName to be %s, got %s", newRoom.Id.String(), got.Id.String())
		}
	})

	t.Run("should get many rooms", func(t *testing.T) {
		mdbq := newMockDBQuery()
		s := NewService(mdbq)
		for i := 0; i < 30; i++ {
			mockDto := dtos.CreateNewRoomDTO{
				UniqueName: faker.Word(),
			}
			s.CreateNewRoom(mockDto)
		}
		mockDto := dtos.GetManyRoomDTO{
			Limit:  10,
			Offset: 2,
		}
		got, gotErr := s.GetManyRooms(mockDto)
		if gotErr != nil {
			t.Errorf("Expect no error, got an error")
		}
		if len(got) != mockDto.Limit {
			t.Errorf("Expected %d, got %d", mockDto.Limit, len(got))
		}
	})

	t.Run("should update one room", func(t *testing.T) {
		mdbq := newMockDBQuery()
		s := NewService(mdbq)

		newRoom, _ := s.CreateNewRoom(dtos.CreateNewRoomDTO{UniqueName: faker.Word()})
		mockDto := dtos.UpdateOneRoomByIDDTO{
			UniqueName: faker.Word(),
		}
		got, gotErr := s.UpdateOneRoomByID(newRoom.Id, mockDto)
		if gotErr != nil {
			t.Errorf("Expect no error, got an error")
		}
		if got.UniqueName != mockDto.UniqueName {
			t.Errorf("Expected %s, got %s", mockDto.UniqueName, got.UniqueName)
		}
	})

	t.Run("should delete one room", func(t *testing.T) {
		mdbq := newMockDBQuery()
		s := NewService(mdbq)

		newRoom, _ := s.CreateNewRoom(dtos.CreateNewRoomDTO{UniqueName: faker.Word()})
		got, gotErr := s.DeleteOneRoomByID(newRoom.Id)
		if gotErr != nil {
			t.Errorf("Expect no error, got an error")
		}
		if got.Id != newRoom.Id {
			t.Errorf("Expected %s, got %s", newRoom.Id, got.Id)
		}
	})
}
