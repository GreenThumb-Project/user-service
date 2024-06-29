package postgres

import (
	"database/sql"
	"time"
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

//------------------------------------------------------------------------------

func (u *UserRepo) CreateUser(user *pb.CreateUsersRequest) (*pb.CreateUsersResponce, error) {
	var users *pb.CreateUsersResponce
	err := u.DB.QueryRow(`
			INSERT INTO 
			users(
				id,
				username,
				email,
				password_hash
			VALUES($1,$2,$3))
		`, user.UserId, user.UserName, user.Email).Scan(&users)

	return &pb.CreateUsersResponce{UserId: users.UserId,
		UserName: users.UserName,
		Email:    users.Email}, err
}

//-------------------------------------------------------------------------------------------

func (u *UserRepo) DeleteUser(id string) (*pb.DeleteUserResponce, error) {

	_, err := u.DB.Exec(`
			UODATE TABLE 
			users 
			SET
				daleted_at=$1 
				where 
					id=$2
			`, time.Now(), id)
	if err != nil {
		return &pb.DeleteUserResponce{Success: false}, err
	}

	return &pb.DeleteUserResponce{Success: true}, nil
}

//------------------------------------------------------------------------------------------

func (u *UserRepo) GetUserByIdProfile(id string) (*pb.GetUserByIdProfileResponces, error) {

	return nil, nil
}

//---------------------------------------------------------------------------------------------

func (u *UserRepo) UpdateUserProfile(userProfil *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponces, error) {
	return nil, nil
}
