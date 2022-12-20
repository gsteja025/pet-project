package services

import (
	"context"
	"fmt"

	ser "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

func (S *Linkedinserver) SearchUser(ctx context.Context, in *pb.SearchRequest) (*pb.Users, error) {
	fmt.Println(in.GetTech().String())
	s1 := ser.Skill{
		Technology: in.GetTech().String(),
	}

	Allskills := []ser.Skill{}
	Allusers := []*pb.User{}

	S.Db.Where(&s1).Find(&Allskills)
	for _, ele := range Allskills {
		Allusers = append(Allusers, &pb.User{Id: uint64(ele.UserID)})
	}
	return &pb.Users{Users: Allusers}, nil
}
