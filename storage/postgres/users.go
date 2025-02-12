package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	pb "user_service/generated/users"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// GET user by id
func (u *UserRepo) GetUserById(id string) (*pb.GetUserByIdResponce, error) {
	user := pb.GetUserByIdResponce{}

	err := u.DB.QueryRow(`
		SELECT
			id,
			username,
			email
		FROM
			users
		WHERE
			deleted_at = 0 and id = $1
	`, id).Scan(&user.UserId, &user.UserName, &user.Email)

	return &user, err
}

//CREATE user

func (u *UserRepo) CreateUser(in *pb.CreateUsersRequest) (*pb.CreateUsersResponce, error) {
	var user pb.CreateUsersResponce
	err := u.DB.QueryRow(`
		INSERT INTO 
		users(
			id,
			username,
			email,
			password
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING 
			id
	`, in.UserId, in.UserName, in.Email, in.Password).Scan(&user.UserId)

	return &pb.CreateUsersResponce{
		UserId: user.UserId}, err
}

// login

func (u *UserRepo) LoginUser(loginReq *pb.LoginRequest) (*pb.LoginResponse, error) {

	userLog := pb.LoginResponse{}
	err := u.DB.QueryRow(`
		SELECT
			id,
			username,
			email
		FROM
			users
		WHERE
			email = $1 and password = $2
	`, loginReq.Email, loginReq.Password).Scan(&userLog.UserId, &userLog.Username, &userLog.Email)

	return &userLog, err
}

// Email exist

func (u *UserRepo) EmailExists(email string) (*pb.EmailExistsResponse, error) {
	var exists bool
	err := u.DB.QueryRow(`
		SELECT
			EXISTS (
				SELECT
					1
				FROM
					users
				WHERE
					email = $1
			)
	`, email).Scan(&exists)

	return &pb.EmailExistsResponse{Exists: exists}, err
}

func (u *UserRepo) UpdateUser(newUser *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res, err := u.DB.Exec(`
		UPDATE 
			users
		SET
			username = $1,
			email = $2,
			password = $3
		WHERE
			user_id = $4
	`, newUser.Username, newUser.Email, newUser.Password, newUser.UserId)

	if err != nil {
		return &pb.UpdateUserResponse{Success: false}, err
	}

	rowsAffect, err := res.RowsAffected()

	if err != nil {
		return &pb.UpdateUserResponse{Success: false}, err
	} else if rowsAffect == 0 {
		return &pb.UpdateUserResponse{Success: false}, errors.New(`user not found `)
	}

	return &pb.UpdateUserResponse{Success: true}, nil
}

//Userni Id si orqali topib delted_at columnni update qilish kerak

func (u *UserRepo) DeleteUser(id string) (*pb.DeleteUserResponce, error) {
	res, err := u.DB.Exec(`
		UPDATE  
			users 
		SET
			deleted_at = date_part('epoch', current_timestamp)::INT 
		WHERE 
			deleted_at = 0 and id=$1
	`, id)

	if err != nil {
		return &pb.DeleteUserResponce{Success: false}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.DeleteUserResponce{Success: false}, err
	}

	return &pb.DeleteUserResponce{Success: true}, nil
}

//GETUserByID profile

func (u *UserRepo) GetUserByIdProfile(id string) (*pb.GetUserByIdProfileResponces, error) {

	var userProfile pb.GetUserByIdProfileResponces

	err := u.DB.QueryRow(`
			SELECT
				up.user_id,
				up.full_name,
				up.bio,
				up.expertise,
				up.location,
				up.avatar_url
			FROM
				users as u
			INNER JOIN
				user_profiles as up ON up.user_id = u.id
			WHERE 
				u.deleted_at = 0 and up.user_id=$1
				`, id).Scan(&userProfile.UserId, &userProfile.FullName, &userProfile.Bio, &userProfile.UserExpertise, &userProfile.Location, &userProfile.AvatarUrl)

	return &userProfile, err
}

func (u *UserRepo) CreateUserProfile(userProfile *pb.CreateProfileUsersRequest) (*pb.CreateProfileUsersResponce, error) {

	_, err := u.DB.Exec(`
		INSERT INTO
		user_profiles(
			user_id,
			full_name,
			bio,
			expertise,
			location,
			avatar_url
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6)
		`, userProfile.UserId, userProfile.FullName, userProfile.Bio,userProfile.UserExpertise, userProfile.Location, userProfile.AvatarUrl)
	

	if err != nil {
		return &pb.CreateProfileUsersResponce{Success: false}, err
	}

	return &pb.CreateProfileUsersResponce{Success: true}, nil
}

func (u *UserRepo) UpdateUserProfile(in *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponces, error) {

	query := ` UPDATE user_profiles SET `
	n := 1
	var arr []interface{}
	if len(in.FullName) > 0 {
		query += fmt.Sprintf(" full_name=$%d, ", n)
		arr = append(arr, in.FullName)
		n++
	}
	if len(in.Bio) > 0 {
		query += fmt.Sprintf(" bio=$%d, ", n)
		arr = append(arr, in.Bio)
		n++
	}
	if len(in.UserExpertise) > 0 {
		query += fmt.Sprintf(" expertise=$%d, ", n)
		arr = append(arr, in.UserExpertise)
		n++
	}
	if len(in.Location) > 0 {
		query += fmt.Sprintf(" location=$%d, ", n)
		arr = append(arr, in.Location)
		n++
	}

	if len(in.AvatarUrl) > 0 {
		query += fmt.Sprintf(" avatar_url=$%d, ", n)
		arr = append(arr, in.AvatarUrl)
		n++
	}
	query = query[:len(query)-2]
	arr = append(arr, in.UserId)
	query += fmt.Sprintf(" where user_id=$%d ", n)

	_, err := u.DB.Exec(query, arr...)

	if err != nil {
		return &pb.UpdateUserProfileResponces{Success: false}, err
	}

	return &pb.UpdateUserProfileResponces{Success: true}, nil

}
