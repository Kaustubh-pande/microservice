package main

import (
	"context"
	"fmt"

	pb "github.com/grpc-master/proto"
	micro "github.com/micro/go-micro"
)

// const (
// 	port = ":50051"
// )

type repository interface {
	Create(*pb.Drug) (*pb.Drug, error)
	GetAll() []*pb.Drug
}

type Repository struct {
	//mu   sync.RWMutex
	Drug []*pb.Drug
}

func (repo *Repository) Create(Drug *pb.Drug) (*pb.Drug, error) {
	//repo.mu.Lock()
	updated := append(repo.Drug, Drug)
	repo.Drug = updated
	//repo.mu.Unlock()
	return Drug, nil
}

type service struct {
	repo repository
}

func (s *service) CreateDrug(ctx context.Context, req *pb.Drug, res *pb.Response) error {

	// Save our consignment
	_, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	//res = &pb.Response{Created: true}
	res.Created = true
	return nil
}
func (repo *Repository) GetAll() []*pb.Drug {
	return repo.Drug
}
func (s *service) GetDrug(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	drug := s.repo.GetAll()
	res.Drug = drug
	return nil
	//&pb.GetResponse{Drug: drug},
}
func main() {

	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("GRPC.MASTER"),
	)
	srv.Init()

	// Set-up our gRPC server.
	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	//pb.RegisterDrugServiceServer(srv.Server(), &service{repo})
	pb.RegisterDrugServiceHandler(srv.Server(), &service{repo})

	// Register reflection service on gRPC server.
	//
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// func allUsers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "All Users Endpoint Hit")
// }

// func newUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "New User Endpoint Hit")
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Delete User Endpoint Hit")
// }

// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Update User Endpoint Hit")
// }
// func handleRequests() {
// 	myRouter := mux.NewRouter() //.StrictSlash(true)
// 	myRouter.HandleFunc("/users", allUsers).Methods("GET")
// 	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
// 	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
// 	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8081", myRouter))
// }

// func main() {
// 	fmt.Println("in main ")

// 	// Handle Subsequent requests
// 	handleRequests()
// }
