package db_query

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/kostmate/common/pkg/entities"
)

type DBQuery struct {
	db *sqlx.DB
}

type DBQuerier interface {
	InsertNewUser(email string, full_name string, role string) (entities.UserEntity, error)
	GetOneUserByID(id uuid.UUID) (entities.UserEntity, error)
	UpdateOneUserByID(id uuid.UUID, email string, full_name string, role string) (entities.UserEntity, error)
	DeleteOneUserByID(id uuid.UUID) (entities.UserEntity, error)
}

func NewDBQuery(db *sqlx.DB) *DBQuery {
	return &DBQuery{db}
}

func (d *DBQuery) InsertNewUser(email string, full_name string, role string) (entities.UserEntity, error) {
	q := `
		insert into public.users (email, full_name, role)
			values ($1, $2, $3) 
			returning id, email, full_name, role, created_at, updated_at;
	`
	stmt, _ := d.db.Preparex(q)

	switch role {
	case "USER":
	case "ADMIN":
		break
	default:
		return entities.UserEntity{}, errors.New("Invalid user role")
	}

	var newUser entities.UserEntity
	if err := stmt.Get(&newUser, email, full_name, role); err != nil {
		return entities.UserEntity{}, errors.New("Failed to create new user")
	}

	return newUser, nil
}

func (d *DBQuery) GetOneUserById(id uuid.UUID) (entities.UserEntity, error) {
	q := `
		select id, email, full_name, role, created_at, updated_at
			where id = $1;
	`
	stmt, _ := d.db.Preparex(q)
	var user entities.UserEntity
	if err := stmt.Get(&user, id); err != nil {
		return entities.UserEntity{}, errors.New("Failed to get user")
	}

	return user, nil
}

func (d *DBQuery) UpdateOneUserById(id uuid.UUID, email string, full_name string, role string) (entities.UserEntity, error) {
	q := `
		update public.users
			set email = $2, full_name = $3, role = $4
			where id = $1
			returning, email, full_name, role, created_at, updated_at;
	`
	stmt, _ := d.db.Preparex(q)

	switch role {
	case "USER":
	case "ADMIN":
		break
	default:
		return entities.UserEntity{}, errors.New("Invalid user role")
	}
	var updatedUser entities.UserEntity
	if err := stmt.Get(&updatedUser, email, full_name, role); err != nil {
		return entities.UserEntity{}, errors.New("Failed to create new user")
	}

	return updatedUser, nil
}

func (d *DBQuery) Delete(id uuid.UUID) (entities.UserEntity, error) {
	q := `
		delete from public.users
			where id
			returning, email, full_name, role, created_at, updated_at;
	`
	stmt, _ := d.db.Preparex(q)

	var deletedUser entities.UserEntity
	if err := stmt.Get(&deletedUser, id); err != nil {
		return entities.UserEntity{}, errors.New("Failed to get user")
	}

	return deletedUser, nil
}
