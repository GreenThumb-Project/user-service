package postgres

import (
	"database/sql"
	pb "user_service/generated/users"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) GetUserById(id string) (*pb.CreateUsersResponce, error) {

	return nil, nil
}

func (u *UserRepo) CreateUser(user *pb.CreateUsersRequest) (*pb.CreateUsersResponce, error) {
	return nil, nil
}


func (u *UserRepo) DeleteUser(id string) (*pb.DeleteUserResponce, error) {
	return nil, nil
}


func (u *UserRepo) GetUserByIdProfile(id string) (*pb.GetUserByIdProfileResponces, error) {
	return nil, nil
}


func (u *UserRepo) UpdateUserProfile(id string) (*pb.UpdateUserProfileResponces, error) {
	return nil, nil
}
