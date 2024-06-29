package main

import (
	"log"
	"net"
	"user_service/generated/users"
	"user_service/service"
	"user_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		log.Fatal("ERROR - ", err.Error())
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatal("ERROR - ", err.Error())
	}

	serv := service.NewUserService(postgres.NewUserRepo(db))
	s := grpc.NewServer()
	users.RegisterUserManagementServer(s, serv)

	log.Println("Server is running on 50050 ...")

	log.Fatal(s.Serve(listener))
}
