package postgres

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	pb "user_service/generated/users"
)

func TestGetUserById(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	user := NewUserRepo(db)

	resp, err := user.GetUserById("c86c3d63-d516-4563-9d68-66724285accb")
	if err != nil {
		panic(err)
	}
	wait := pb.GetUserByIdResponce{
		UserId:   "c86c3d63-d516-4563-9d68-66724285accb",
		UserName: "user3",
		Email:    "user3@example.com",
	}

	if !reflect.DeepEqual(resp, &wait) {
		t.Errorf("have %v , wont %v", resp, &wait)
	}

}

func TestCreateUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	user := NewUserRepo(db)

	respcreate, _ := user.CreateUser(&pb.CreateUsersRequest{
		UserId:   "65d52c55-5641-4f4c-b98a-28701deadf5f",
		UserName: "userq123",
		Email:    "user1@exampsdfdjdsle.come2",
		Password: "qwerty2004",
	})

	waitcreate := pb.CreateUsersResponce{
		UserId:   "65d52c55-5641-4f4c-b98a-28701deadf5f",
	}

	if reflect.DeepEqual(respcreate, &waitcreate) {
		t.Errorf("have %v , wont %v", respcreate, &waitcreate)
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	user := NewUserRepo(db)

	respcreate, _ := user.DeleteUser("0415e108-cdba-4f6d-9e44-e9268152e23f")

	waitcreate := pb.DeleteUserResponce{
		Success: true,
	}

	if !reflect.DeepEqual(respcreate, &waitcreate) {
		t.Errorf("have %v , wont %v", respcreate, &waitcreate)
	} 
}

func TestGetUserByIdProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	user := NewUserRepo(db)

	respcreate, err := user.GetUserByIdProfile("c86c3d63-d516-4563-9d68-66724285accb")
	if err != nil {
		log.Fatal(err)
	}
	waitcreate := pb.GetUserByIdProfileResponces{
		UserId:        "c86c3d63-d516-4563-9d68-66724285accb",
		FullName:      "User Three",
		Bio:           "Bio for User Three",
		UserExpertise: "beginner",
		Location:      "Location Three",
		AvatarUrl:     "http://example.com/avatar3.png",
	}

	if !reflect.DeepEqual(respcreate, &waitcreate) {
		t.Errorf("have %v , wont %v", respcreate, &waitcreate)
	}
}

func TestCreateUserProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	user := NewUserRepo(db)
	
	respcreate, err := user.CreateUserProfile(&pb.CreateProfileUsersRequest{
		UserId:        "c16c9fdc-3c06-46c1-a5a6-d1acc89bc101",
		FullName:      "dnjokofe",
		Bio:           "jsdhokud",
		UserExpertise: "beginner",
		Location:      "sjdfkoijfuu",
		AvatarUrl:     "dfngijkmien",
	})
	if err != nil {
		fmt.Println("sjkojijolkjiisdic")
		log.Fatal(err)
	}
	waitcreate := pb.CreateProfileUsersResponce{
		Success: true,
	}

	if !reflect.DeepEqual(respcreate, &waitcreate) {
		t.Errorf("have %v , wont %v", respcreate, &waitcreate)
	}
}

func TestUpdateUserProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	user := NewUserRepo(db)

	respcreate, err := user.UpdateUserProfile(&pb.UpdateUserProfileRequest{UserId: "c86c3d63-d516-4563-9d68-66724285accb",
		FullName: "edjfsdfff"})

	if err != nil {
		panic(err)
	}

	waitcreate := pb.UpdateUserProfileResponces{
		Success: true,
	}

	if !reflect.DeepEqual(respcreate, &waitcreate) {
		t.Errorf("have %v , wont %v", respcreate, &waitcreate)
	}
}
