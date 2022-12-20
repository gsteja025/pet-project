package services

import (
	"context"
	"log"
	"reflect"

	"example.com/petproject/Email"
	ser "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

func (s *Linkedinserver) GetConnectedUsers(ctx context.Context, in *pb.User) (*pb.Users, error) {
	log.Printf("Connected users")
	connects := []ser.Connected{}
	connects1 := []ser.Connected{}
	Finalconnects := []*pb.User{}
	// str1 := "Connected"
	s.Db.Where(&ser.Connected{
		User_1: uint(in.GetId()),
		Status: "Connected",
	}).Find(&connects)
	// s.Db.Where("user_2 = ? AND status = ?", in.GetId(), str1).Find(&connects1)
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
	str := "pending"
	var conn ser.Connected
	conn1 := conn
	usermodel := ser.User{}
	s.Db.Where("user_1 = ? and user_2 = ?", in.GetId2(), in.GetId1()).Find(&conn)
	if !reflect.DeepEqual(conn, conn1) && conn.Status == "pending" {
		str = "Connected"
		s.Db.Save(&ser.Connected{User_1: uint(in.GetId2()), User_2: uint(in.GetId1()), Status: str})
		s.Db.Where("id = ?", in.GetId2()).Find(&usermodel)
		Email.SendEmail(usermodel.Email)

	} else {
		s.Db.Save(&ser.Connected{User_1: uint(in.GetId1()), User_2: uint(in.GetId2()), Status: str})
	}
	return &pb.Emptyresponse{}, nil
}
