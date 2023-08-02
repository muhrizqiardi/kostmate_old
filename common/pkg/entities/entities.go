package entities

import "github.com/google/uuid"

type UserEntity struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FullName  string    `json:"fullName" db:"full_name"`
	Role      string    `json:"role" db:"role"`
	CreatedAt string    `json:"createdAt" db:"created_at"`
	UpdatedAt string    `json:"updatedAt" db:"updated_at"`
}

type RoomEntity struct {
	Id         uuid.UUID `json:"id" db:"id"`
	UniqueName string    `json:"unqiueName" db:"unique_name"`
	CreatedAt  string    `json:"createdAt" db:"created_at"`
	UpdatedAt  string    `json:"updatedAt" db:"updated_at"`
}

type PaymentEntity struct {
	Id        uuid.UUID `json:"id" db:"id"`
	CreatedAt string    `json:"createdAt" db:"created_at"`
	UpdatedAt string    `json:"updatedAt" db:"updated_at"`
}

type BookingEntity struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Status    string    `json:"status" db:"status"`
	CreatedAt string    `json:"createdAt" db:"created_at"`
	UpdatedAt string    `json:"updatedAt" db:"updated_at"`
}
