package postgres

import (
	"database/sql"
	"fmt"
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
			UPDATE TABLE 
			users 
			SET
				deleted_at = date_part('epoch', current_timestamp)::INT 
				where 
					id=$1
			`, id)
	if err != nil {
		return &pb.DeleteUserResponce{Success: false}, err
	}

	return &pb.DeleteUserResponce{Success: true}, nil
}

//------------------------------------------------------------------------------------------

func (u *UserRepo) GetUserByIdProfile(id string) (*pb.GetUserByIdProfileResponces, error) {

	var userProfile *pb.GetUserByIdProfileResponces

	err := u.DB.QueryRow(`
			SELECT
				user_id,
				full_name,
				bio,
				user_expertise,
				location,
				avatar_url
			FROM
				user_profiles
			WHERE user_id=$1
				`, id).Scan(&userProfile)

	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

//---------------------------------------------------------------------------------------------

func (u *UserRepo) UpdateUserProfile(in *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponces, error) {

	query := `update user_profiles set `
	n := 1
	var arr []interface{}
	if len(in.FulName) > 0 {
		query += fmt.Sprintf("full_name=$%d, ", n)
		arr = append(arr, in.FulName)
		n++
	}
	if len(in.Bio) > 0 {
		query += fmt.Sprintf("bio=$%d, ", n)
		arr = append(arr, in.Bio)
		n++
	}
	if len(in.UserExpertise) > 0 {
		query += fmt.Sprintf("expertise=$%d, ", n)
		arr = append(arr, in.UserExpertise)
		n++
	}
	if len(in.Location) > 0 {
		query += fmt.Sprintf("location=$%d, ", n)
		arr = append(arr, in.Location)
		n++
	}

	if len(in.AvataEUrl) > 0 {
		query += fmt.Sprintf("avatar_url=$%d, ", n)
		arr = append(arr, in.AvataEUrl)
		n++
	}
	arr = append(arr, in.UserId)

	query += fmt.Sprintf(" where user_id=$%d ", n)

	_, err := u.DB.Exec(query, arr...)
	
	if err != nil {
		return &pb.UpdateUserProfileResponces{Success: false}, err
	}

	return &pb.UpdateUserProfileResponces{Success: true}, nil

}
