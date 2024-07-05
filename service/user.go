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

func (s *UserService) CheckUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.User.GetUserById(in.UserId)

	if err != nil {
		return &pb.UserResponse{Exists: false}, err
	}

	exists := user.UserId == in.UserId

	return &pb.UserResponse{Exists: exists}, nil
}

func (s *UserService) EmailExists(ctx context.Context, in *pb.EmailExistsRequest) (*pb.EmailExistsResponse, error) {
	return s.User.EmailExists(in.Email)
}

func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.User.LoginUser(in)
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

	return resp, err
}

func (s UserService) GetUserByIdProfile(ctx context.Context, req *pb.GetUserByIdProfileRequest) (*pb.GetUserByIdProfileResponces, error) {
	return s.User.GetUserByIdProfile(req.UserId)
}

func (s UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponces, error) {
	return s.User.UpdateUserProfile(req)
}
