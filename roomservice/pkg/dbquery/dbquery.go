package dbquery

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/kostmate/common/pkg/entities"
)

const (
	QueryInsertNewRoom = `
		insert into public.rooms (uniqueName)
			values ($1) 
			returning id, unique_name, created_at, updated_at;
	`
	QueryGetOneRoomByID = `
		select id, unique_name, created_at, updated_at
			from public.rooms
			where id = $1;
	`
	QueryGetOneRoomByUniqueName = `
		select id, unique_name, created_at, updated_at
			from public.rooms
			where unique_name = $1;
	`
	QueryGetManyRooms = `
		select id, unique_name, created_at, updated_at
			from public.rooms
			limit $1 offset $2;
	`
	QueryUpdateOneRoomByID = `
		update public.rooms
			set unique_name = $2
			where id = $1
			returning id, unique_name, created_at, updated_at;
	`
	QueryDeleteOneRoomByID = `
		delete from public.rooms
			where id = $1
			returning id, unique_name, created_at, updated_at;
	`
)

type DBQuery struct {
	db *sqlx.DB
}

type DBQuerier interface {
	InsertNewRoom(uniqueName string) (entities.RoomEntity, error)
	GetOneRoomByID(id uuid.UUID) (entities.RoomEntity, error)
	GetOneRoomByUniqueName(uniqueName string) (entities.RoomEntity, error)
	GetManyRooms(limit int, offset int) ([]entities.RoomEntity, error)
	UpdateOneRoomByID(id uuid.UUID, uniqueName string) (entities.RoomEntity, error)
	DeleteOneRoomByID(id uuid.UUID) (entities.RoomEntity, error)
}

func NewDBQuery(db *sqlx.DB) *DBQuery {
	return &DBQuery{db}
}

func (dbq *DBQuery) InsertNewRoom(uniqueName string) (entities.RoomEntity, error) {
	stmt, _ := dbq.db.Preparex(QueryInsertNewRoom)

	var newRoom entities.RoomEntity
	if err := stmt.Get(&newRoom, uniqueName); err != nil {
		return entities.RoomEntity{}, errors.New("Failed to create new room")
	}

	return newRoom, nil
}

func (dbq *DBQuery) GetOneRoomByID(id uuid.UUID) (entities.RoomEntity, error) {
	stmt, _ := dbq.db.Preparex(QueryGetOneRoomByID)

	var room entities.RoomEntity
	if err := stmt.Get(&room, id); err != nil {
		return entities.RoomEntity{}, errors.New("Failed to create new room")
	}

	return room, nil
}

func (dbq *DBQuery) GetOneRoomByUniqueName(uniqueName string) (entities.RoomEntity, error) {
	stmt, _ := dbq.db.Preparex(QueryGetOneRoomByUniqueName)

	var room entities.RoomEntity
	if err := stmt.Get(&room, uniqueName); err != nil {
		return entities.RoomEntity{}, errors.New("Failed to create new room")
	}

	return room, nil
}

func (dbq *DBQuery) GetManyRooms(limit int, offset int) ([]entities.RoomEntity, error) {
	stmt, _ := dbq.db.Preparex(QueryGetManyRooms)

	var room []entities.RoomEntity
	if err := stmt.Select(&room, limit, offset); err != nil {
		return []entities.RoomEntity{}, errors.New("Failed to create new room")
	}

	return room, nil
}

func (dbq *DBQuery) UpdateOneRoomByID(id uuid.UUID, uniqueName string) (entities.RoomEntity, error) {
	stmt, _ := dbq.db.Preparex(QueryUpdateOneRoomByID)

	var updatedRoom entities.RoomEntity
	if err := stmt.Get(&updatedRoom, uniqueName); err != nil {
		return entities.RoomEntity{}, errors.New("Failed to create new room")
	}

	return updatedRoom, nil
}

func (dbq *DBQuery) DeleteOneRoomByID(id uuid.UUID) (entities.RoomEntity, error) {
	stmt, _ := dbq.db.Preparex(QueryDeleteOneRoomByID)

	var deletedRoom entities.RoomEntity
	if err := stmt.Get(&deletedRoom, id); err != nil {
		return entities.RoomEntity{}, errors.New("Failed to create new room")
	}

	return deletedRoom, nil
}
