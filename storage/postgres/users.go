package postgres

import (
	"database/sql"
	u "user_service/generated/users"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) GetUserById(id string) (*u.CreateUsersResponce, error) {

	return nil, nil
}	
