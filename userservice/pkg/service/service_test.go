package service

import (
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/muhrizqiardi/kostmate/common/pkg/entities"
	"github.com/muhrizqiardi/kostmate/userservice/pkg/dtos"
)

type mockDBQuery struct {
}

func newMockDBQuery() *mockDBQuery {
	return &mockDBQuery{}
}

func (mdbq *mockDBQuery) InsertNewUser(email string, full_name string, role string) (entities.UserEntity, error) {
	newUUID, _ := uuid.NewRandom()

	return entities.UserEntity{
		Id:        newUUID,
		Email:     email,
		FullName:  full_name,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (mdbq *mockDBQuery) GetOneUserByID(id uuid.UUID) (entities.UserEntity, error) {
	return entities.UserEntity{
		Id:        id,
		Email:     faker.Email(),
		FullName:  faker.FirstName() + " " + faker.LastName(),
		Role:      "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (mdbq *mockDBQuery) UpdateOneUserByID(id uuid.UUID, email string, full_name string, role string) (entities.UserEntity, error) {
	return entities.UserEntity{
		Id:        id,
		Email:     faker.Email(),
		FullName:  faker.FirstName() + " " + faker.LastName(),
		Role:      "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (mdbq *mockDBQuery) DeleteOneUserByID(id uuid.UUID) (entities.UserEntity, error) {
	return entities.UserEntity{
		Id:        id,
		Email:     faker.Email(),
		FullName:  faker.FirstName() + " " + faker.LastName(),
		Role:      "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func TestService_CreateNewUser(t *testing.T) {
	t.Run("should create new user", func(t *testing.T) {
		mockPayload_userRole := dtos.CreateNewUserDTO{
			Email:    faker.Email(),
			FullName: faker.FirstName() + " " + faker.LastName(),
			Role:     "USER",
		}

		mockdbq := newMockDBQuery()
		s := NewService(mockdbq)
		got, gotErr := s.CreateNewUser(mockPayload_userRole)
		if gotErr != nil {
			t.Errorf("Expected no error, got error")
		}
		if got.Email != mockPayload_userRole.Email ||
			got.FullName != mockPayload_userRole.FullName ||
			got.Role != mockPayload_userRole.Role {
			t.Errorf("Expected correct email, fullname, or role, got incorrect")
		}
	})
}

func TestService_GetOneUserByID(t *testing.T) {
	t.Run("should get one user by ID", func(t *testing.T) {
		mockdbq := newMockDBQuery()
		s := NewService(mockdbq)
		mockUserId, _ := uuid.NewUUID()
		mockUser := entities.UserEntity{
			Id:        mockUserId,
			Email:     faker.Email(),
			FullName:  faker.FirstName() + " " + faker.LastName(),
			Role:      "USER",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		got, gotErr := s.GetOneUserByID(mockUserId)
		if gotErr != nil {
			t.Errorf("Expected no error, got error")
		}
		if got.Id != mockUserId {
			t.Errorf("Expected %s, got %s", mockUser.Id, got.Id)
		}
	})
}

func TestService_UpdateOneUserByID(t *testing.T) {
	t.Run("should update one user by ID", func(t *testing.T) {
		mockdbq := newMockDBQuery()
		s := NewService(mockdbq)
		mockUserId, _ := uuid.NewUUID()
		mockUser := entities.UserEntity{
			Id:        mockUserId,
			Email:     faker.Email(),
			FullName:  faker.FirstName() + " " + faker.LastName(),
			Role:      "USER",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		got, gotErr := s.UpdateOneUserByID(mockUserId, dtos.UpdateOneUserByIDDTO{
			Email:    mockUser.Email,
			FullName: mockUser.FullName,
			Role:     mockUser.Role,
		})
		if gotErr != nil {
			t.Errorf("Expected no error, got error")
		}
		if got.Id != mockUserId {
			t.Errorf("Expected %s, got %s", mockUser.Id, got.Id)
		}
	})
}

func TestService_DeleteOneUserByID(t *testing.T) {
	t.Run("should get one user by ID", func(t *testing.T) {
		mockdbq := newMockDBQuery()
		s := NewService(mockdbq)
		mockUserId, _ := uuid.NewUUID()
		mockUser := entities.UserEntity{
			Id:        mockUserId,
			Email:     faker.Email(),
			FullName:  faker.FirstName() + " " + faker.LastName(),
			Role:      "USER",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		got, gotErr := s.GetOneUserByID(mockUserId)
		if gotErr != nil {
			t.Errorf("Expected no error, got error")
		}
		if got.Id != mockUserId {
			t.Errorf("Expected %s, got %s", mockUser.Id, got.Id)
		}
	})
}
