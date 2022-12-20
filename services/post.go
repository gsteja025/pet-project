package services

import (
	"context"
	"log"

	ser "example.com/petproject/models"
	pb "example.com/petproject/protos"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Linkedinserver struct {
	pb.UnimplementedLinkedinDatabaseCrudServer
	Db *gorm.DB
}

func (s *Linkedinserver) Createpost(ctx context.Context, in *pb.NewPost) (*pb.Post, error) {
	log.Printf("creating new post called")
	newpos := ser.Post{
		Text:   in.GetText(),
		UserID: 1,
	}
	s.Db.Save(&newpos)
	return &pb.Post{Text: in.GetText(), Id: uint64(newpos.ID)}, nil
}

func (s *Linkedinserver) GetPostComments(ctx context.Context, in *pb.Post) (*pb.Comments, error) {
	log.Printf("Getting comments of post")
	allcommen := []ser.Comment{}
	Finalcomments := []*pb.Comment{}
	s.Db.Where("post_id = ?", in.GetId()).Find(&allcommen)
	for _, conn := range allcommen {
		Finalcomments = append(Finalcomments, &pb.Comment{Id: uint64(conn.CommentID), Text: conn.Text, Commenterid: uint64(conn.CommenterId)})
	}

	return &pb.Comments{Allcomments: Finalcomments}, nil
}

func (s *Linkedinserver) GetPostLikes(ctx context.Context, in *pb.Post) (*pb.Users, error) {
	log.Printf("Getting likes of post")
	allLikes := []ser.Likes{}
	FinalLikes := []*pb.User{}
	s.Db.Where("post_id = ?", in.GetId()).Find(&allLikes)
	for _, conn := range allLikes {
		FinalLikes = append(FinalLikes, &pb.User{Id: uint64(conn.LikerId)})
	}
	return &pb.Users{Users: FinalLikes}, nil
}

func (s *Linkedinserver) LikeOtherPost(ctx context.Context, in *pb.Post) (*pb.Emptyresponse, error) {
	posts := ser.Likes{
		PostID:  uint(in.GetId()),
		LikerId: uint(in.UserID),
	}
	s.Db.Save(&posts)
	return &pb.Emptyresponse{}, nil
}
