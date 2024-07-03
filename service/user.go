package service

import (
	"context"
	pb "user_service/generated/users"
	"user_service/storage/postgres"
)

type UserService struct {
	pb.UnimplementedUserManagementServer
	User *postgres.UserRepo
}

func NewUserService(user *postgres.UserRepo) *UserService {
	return &UserService{User: user}
}

func (s UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponce, error) {
	return s.User.GetUserById(req.UserId)
}

func (s UserService) CreateUser(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersResponce, error) {
	return s.User.CreateUser(req)
}

func (s UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponce, error) {
	return s.User.DeleteUser(req.UserId)
}

func (s *UserService) CreateUserProfile(ctx context.Context, in *pb.CreateProfileUsersRequest) (*pb.CreateProfileUsersResponce, error) {
	resp, err := s.User.CreateUserProfile(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s UserService) GetUserByIdProfile(ctx context.Context, req *pb.GetUserByIdProfileRequest) (*pb.GetUserByIdProfileResponces, error) {
	return s.User.GetUserByIdProfile(req.UserId)
}

func (s UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponces, error) {
	return s.User.UpdateUserProfile(req)
}
