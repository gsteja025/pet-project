package services

import (
	"context"
	"log"
	"reflect"

	Email "example.com/petproject/Email"
	ser "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

func (s Linkedinserver) GetConnectedUsers(ctx context.Context, in *pb.User) (*pb.Users, error) {
	log.Printf("Connected users")
	connects := []ser.Connected{}
	connects1 := []ser.Connected{}
	Finalconnects := []*pb.User{}
	// str1 := "Connected"
	connects, err := s.Db.GetConnectedUsersDbInteraction(ser.Connected{User_1: uint(in.GetId())})
	// s.Db.Where("user_2 = ? AND status = ?", in.GetId(), str1).Find(&connects1)
	if err != nil {
		panic(err.Error())
	}
	for _, conn := range connects {
		Finalconnects = append(Finalconnects, &pb.User{Id: uint64(conn.User_2)})
	}

	for _, conn := range connects1 {
		Finalconnects = append(Finalconnects, &pb.User{Id: uint64(conn.User_1)})
	}
	return &pb.Users{Users: Finalconnects}, nil

}

func (s *Linkedinserver) ConnectWithOtherUser(ctx context.Context, in *pb.ConnectionRequest) (*pb.Emptyresponse, error) {
	log.Printf("YOU NOW HAVE A NEW CONNECTION")
	// var allconnectsids []uint64
	// cnt := 0
	var conn ser.Connected
	conn1 := conn
	usermodel := ser.User{}
	Userslice := []ser.User{}
	user1 := ser.User{}
	user2 := ser.User{}
	user1.ID = uint(in.GetId1())
	user2.ID = uint(in.GetId2())
	Userslice = append(Userslice, user1)
	Userslice = append(Userslice, user2)
	conn, err := s.Db.ConnectWithOtherUserDbinteraction1(Userslice)
	if err != nil {
		panic(err.Error())
	}
	if !reflect.DeepEqual(conn, conn1) && conn.Status == "pending" {
		s.Db.ConnectWithOtherUserDbinteraction2(Userslice)
		Email.SendEmail(usermodel.Email)

	} else {
		s.Db.ConnectWithOtherUserDbinteraction3(Userslice)
	}
	return &pb.Emptyresponse{}, nil
}
