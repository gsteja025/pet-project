package services

import (
	"context"
	"log"

	ser "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

func (s Linkedinserver) GetConnectedUsers(ctx context.Context, in *pb.User) (*pb.Users, error) {
	log.Printf("Connected users")
	connects := []ser.Connected{}
	Finalconnects := []*pb.User{}
	// str1 := "Connected"
	connects, err := s.Db.GetConnectedUsersDbInteraction(ser.Connected{User_1: uint(in.GetId())})
	// s.Db.Where("user_2 = ? AND status = ?", in.GetId(), str1).Find(&connects1)
	if err != nil {
		panic(err.Error())
	}
	for _, conn := range connects {
		if conn.User_1 == uint(in.GetId()) {
			Finalconnects = append(Finalconnects, &pb.User{Id: uint64(conn.User_2)})
		} else {
			Finalconnects = append(Finalconnects, &pb.User{Id: uint64(conn.User_1)})
		}
	}

	return &pb.Users{Users: Finalconnects}, nil

}

func (s *Linkedinserver) ConnectWithOtherUser(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	log.Printf("YOU NOW HAVE A NEW CONNECTION")
	// var allconnectsids []uint64
	// cnt := 0
	// usermodel := ser.User{}
	user1 := ser.User{}
	user2 := ser.User{}
	user1.ID = uint(in.GetId1())
	user2.ID = uint(in.GetId2())
	var Userslice = []ser.User{user1, user2}
	conn, err := s.Db.ConnectWithOtherUser(Userslice)
	return &pb.ConnectionResponse{Message: conn.Status}, err
}
