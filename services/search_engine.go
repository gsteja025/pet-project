package services

import (
	"context"
	"fmt"

	model "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

func (S *Linkedinserver) SearchUser(ctx context.Context, in *pb.SearchRequest) (*pb.Users, error) {
	fmt.Println(in.GetTech().String())
	s1 := model.Skill{
		Technology: in.GetTech().String(),
	}

	Allskills := []model.Skill{}
	Allusers := []*pb.User{}

	Allskills, err := S.Db.SearchUserDbinteraction(s1)
	if err != nil {
		fmt.Printf("Thers an error")
	}
	for _, ele := range Allskills {
		Allusers = append(Allusers, &pb.User{Id: uint64(ele.UserID)})
	}
	return &pb.Users{Users: Allusers}, err
}
