package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FullName  string    `json:"fullName" db:"full_name"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type RoomEntity struct {
	Id         uuid.UUID `json:"id" db:"id"`
	UniqueName string    `json:"unqiueName" db:"unique_name"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

type PaymentEntity struct {
	Id        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type BookingEntity struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
