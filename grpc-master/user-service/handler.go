package main

import (
	"golang.org/x/net/context"
	pb "github.com/grpc-master/user-service/proto"

)
type service struct{
	repo repository
	tokenService Authable
}
func(srv *service) Get(ctx context.Context,Req *pb.User,res *pb.Response) error{
	user,err:=srv.repo.Get(req.Id)
	if err =nil{
		return err
	}
	res.User = user
	return nil
}
func (srv *service) GetAll(ctx context.Context,req *pb.User,res *pb.Response)error{
	users ,err := srv.repo.GetAll()
	if err = nil{
		return err
	} 
	res.User = users
	return nil
}
func (srv *service) Auth(ctx context.Context,req *pb.User,res *pb.Response) error {
	user, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingabc"
	return nil
}
func(srv *service) ValidateToken(ctx context.Context,req *pb.Token,res *pb.Token)error{
	return nil
}
func (srv *service) Create(ctx context.Context,req *pb.User,res *pb.Response) error {
	if err :=srv.Create(req);err !=nil{
		return nil
	}
	res.User = req
	return nil
}