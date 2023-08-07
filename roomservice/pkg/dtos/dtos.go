package dtos

type CreateNewRoomDTO struct {
	UniqueName string `json:"unqiueName" db:"unique_name"`
}

type GetManyRoomDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type UpdateOneRoomByIDDTO struct {
	UniqueName string `json:"unqiueName" db:"unique_name"`
}
