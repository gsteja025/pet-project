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
		UserID: uint(in.UserID),
	}

	ans, err := s.Db.CreatepostDbInteraction(newpos)

	return &pb.Post{Id: uint64(ans.ID)}, err
}

func (s *Linkedinserver) GetPostComments(in *pb.PostRequest, stream pb.LinkedinDatabaseCrud_GetPostCommentsServer) error {
	log.Printf("Getting comments of post")
	// Finalcomments := []*pb.Comment{}
	pos := ser.Post{}
	pos.ID = uint(in.GetId())
	comm, err := s.Db.GetPostCommentsDbinteraction(pos)

	if err != nil {
		fmt.Printf("Thers an error")
	}
	// s.Db.Where("post_id = ?", in.GetId()).Find(&allcommen)
	for _, conn := range comm {
		ele1 := pb.Comment{
			Text:   conn.Text,
			PostID: uint64(conn.PostID),
		}
		err := stream.Send(&ele1)
		if err != nil {
			fmt.Println("there is error in postcomments function")
		}

		//Finalcomments = append(Finalcomments, &pb.Comment{Id: uint64(conn.CommentID), Text: conn.Text, Commenterid: uint64(conn.CommenterId)})
	}

	return nil
}

func (s *Linkedinserver) GetPostLikes(ctx context.Context, in *pb.PostRequest) (*pb.Users, error) {
	log.Printf("Getting likes of post")
	allLikes := []ser.Like{}
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
	return &pb.Users{Users: FinalLikes}, err
}

func (s *Linkedinserver) LikeOtherPost(ctx context.Context, in *pb.Request) (*pb.ConnectionResponse, error) {
	posts := ser.Like{
		PostID:  uint(in.GetPostID()),
		LikerId: uint(in.LikerID),
	}
	var message string = "No error found"
	err := s.Db.LikeOtherPostDbinteraction(posts)
	if err != nil {
		fmt.Printf("Thers an error")
		message = "Error found"
	}
	return &pb.ConnectionResponse{Message: message}, err
}
