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
			username,
			email
		)
		VALUES($1,$2,$3)
		RETURNING 
			id,
			username,
			email
	`, in.UserId, in.UserName, in.Email).Scan(&user.UserId, &user.UserName, &user.Email)

	fmt.Println(user.Email)
	return &pb.CreateUsersResponce{
		UserId:   user.UserId,
		UserName: user.UserName,
		Email:    user.Email}, err
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

	if err != nil{
		return &pb.DeleteUserResponce{Success: false}, err
	}
		
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0{
		return &pb.DeleteUserResponce{Success: false}, err
	}

	return &pb.DeleteUserResponce{Success: true}, nil
}

//GETUserByID profile 

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
				user_profiles up
			INNER JOIN
				users u ON up.user_id = u.user_id
			WHERE 
				u.deleted_at = 0 and user_id=$1
				`, id).Scan(&userProfile)

	return userProfile, err
}




func (u *UserRepo) CreateUserProfile(userProfile *pb.CreateProfileUsersRequest) (*pb.CreateProfileUsersResponce, error) {

	_, err := u.DB.Exec(`
		INSERT INTO
			user_id,
			full_name,
			bio,
			user_expertise,
			location,
			avatar_url
		FROM
			user_profiles up
		`, userProfile.UserId, userProfile.FullName, userProfile.Bio, 
		userProfile.UserExpertise, userProfile.Location, userProfile.AvatarUrl)

	
	if err !=  nil {
		return &pb.CreateProfileUsersResponce{Success: false}, err
	}

	return &pb.CreateProfileUsersResponce{Success: true}, nil
}

//Update user profile
func (u *UserRepo) UpdateUserProfile(in *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponces, error) {
	var checker int
	err := u.DB.QueryRow(`
		SELECT 
			deleted_at 
		FROM 
			users 
		WHERE 
			deleted_ad = 0 and id=$1
		`, in.UserId).Scan(&checker)
	if checker != 0 || err != nil {
		return nil, fmt.Errorf("%s user is not found or already deleted", in.UserId)
	}

	query := `UPDATE user_profiles SET `
	n := 1
	var arr []interface{}
	if len(in.FullName) > 0 {
		query += fmt.Sprintf("full_name=$%d, ", n)
		arr = append(arr, in.FullName)
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

	if len(in.AvatarUrl) > 0 {
		query += fmt.Sprintf("avatar_url=$%d, ", n)
		arr = append(arr, in.AvatarUrl)
		n++
	}
	arr = append(arr, in.UserId)

	query += fmt.Sprintf(" WHERE user_id=$%d ", n)

	_, err = u.DB.Exec(query, arr...)

	if err != nil {
		return &pb.UpdateUserProfileResponces{Success: false}, err
	}

	return &pb.UpdateUserProfileResponces{Success: true}, nil

}
