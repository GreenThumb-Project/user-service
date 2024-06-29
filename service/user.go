package service

import (
	"context"
	pb "user_service/generated/users"
	"user_service/storage/postgres"
)

type ServerUser struct {
	pb.UnimplementedUserManagementServer
	DB *postgres.UserRepo
}

func NewUserService(db *postgres.UserRepo) *ServerUser {
	return &ServerUser{DB: db}
}

func (s ServerUser) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponce, error) {
	return s.DB.GetUserById(req.UserId)
}

func (s ServerUser) CreateUser(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersResponce, error) {
	return s.DB.CreateUser(req)
}

func (s ServerUser) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponce, error) {
	return s.DB.DeleteUser(req.UserId)
}

func (s ServerUser) GetUserByIdProfile(ctx context.Context, req *pb.GetUserByIdProfileRequest) (*pb.GetUserByIdProfileResponces, error) {
	return s.DB.GetUserByIdProfile(req.UserId)
}

func (s ServerUser) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponces, error) {
	return s.DB.UpdateUserProfile(req)
}
