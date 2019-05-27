package main

import (
	"log"

	pb "github.com/grpc-master/user-service/proto"
	"github.com/micro/go-micro"
	//_ "github.com/micro/go-plugins/registry/mdns"
	//_ "github.com/micro/go-micro/tree/master/registry/mdns"
	//k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenservice := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("user.service"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterUserserviceHandler(srv.server(), &service{repo, tokenService})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
