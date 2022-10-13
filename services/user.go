package services

import (
	"context"
	"fmt"
	"testing-grpc-with-go/pb"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//}

func NewUserService() *UserService{
	return &UserService{}
}

type UserService struct{
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error){

	//Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id: "123",
		Name: req.GetName(),
		Email: req.GetEmail(),
	}, nil

}
