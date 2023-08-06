package dtos

type CreateNewUserDTO struct {
	Email    string `json:"email" db:"email"`
	FullName string `json:"fullName" db:"full_name"`
	Role     string `json:"role" db:"role"`
}

type UpdateOneUserByIDDTO struct {
	Email    string `json:"email" db:"email"`
	FullName string `json:"fullName" db:"full_name"`
	Role     string `json:"role" db:"role"`
}
