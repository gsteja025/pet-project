package services

import (
	"context"
	"fmt"
	"log"

	"example.com/petproject/database"
	ser "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

type Linkedinserver struct {
	pb.UnimplementedLinkedinDatabaseCrudServer
	Db database.Dbinterface
}

func (s *Linkedinserver) Createpost(ctx context.Context, in *pb.NewPost) (*pb.Post, error) {
	log.Printf("creating new post called")
	newpos := ser.Post{
		Text:   in.GetText(),
		UserID: 1,
	}
	s.Db.CreatepostDbInteraction(newpos)
	return &pb.Post{Text: in.GetText(), UserID: in.GetUserID()}, nil
}

func (s *Linkedinserver) GetPostComments(ctx context.Context, in *pb.Post) (*pb.Comments, error) {
	log.Printf("Getting comments of post")
	Finalcomments := []*pb.Comment{}
	pos := ser.Post{}
	pos.ID = uint(in.GetId())
	comm, err := s.Db.GetPostCommentsDbinteraction(pos)

	if err != nil {
		fmt.Printf("Thers an error")
	}

	// s.Db.Where("post_id = ?", in.GetId()).Find(&allcommen)
	for _, conn := range comm {
		Finalcomments = append(Finalcomments, &pb.Comment{Id: uint64(conn.CommentID), Text: conn.Text, Commenterid: uint64(conn.CommenterId)})
	}

	return &pb.Comments{Allcomments: Finalcomments}, nil
}

func (s *Linkedinserver) GetPostLikes(ctx context.Context, in *pb.Post) (*pb.Users, error) {
	log.Printf("Getting likes of post")
	allLikes := []ser.Likes{}
	FinalLikes := []*pb.User{}
	post := ser.Post{}
	post.ID = uint(in.GetId())
	allLikes, err := s.Db.GetPostLikesDbinteraction(post)
	if err != nil {
		fmt.Printf("Thers an error")
	}
	for _, conn := range allLikes {
		FinalLikes = append(FinalLikes, &pb.User{Id: uint64(conn.LikerId)})
	}
	return &pb.Users{Users: FinalLikes}, nil
}

func (s *Linkedinserver) LikeOtherPost(ctx context.Context, in *pb.Request) (*pb.Emptyresponse, error) {
	posts := ser.Likes{
		PostID:  uint(in.GetPostID()),
		LikerId: uint(in.LikerID),
	}
	err := s.Db.LikeOtherPostDbinteraction(posts)
	if err != nil {
		fmt.Printf("Thers an error")
	}
	return &pb.Emptyresponse{}, nil
}
