package services

import (
	"context"
	"log"

	database "example.com/petproject/database"
	mod "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

type Server struct {
	Db database.Dbinterface
}

func (s Server) CreateComment(ctx context.Context, in *pb.Comment) (*pb.Comment, error) {
	log.Printf("creating new Comment")
	newcom := mod.Comment{
		Text:        in.GetText(),
		CommenterId: uint(in.GetCommenterid()),
		PostID:      uint(in.GetPostID()),
	}
	s.Db.CreateCommentDbInteraction(newcom)
	return &pb.Comment{Id: uint64(newcom.CommentID), Text: in.GetText(), Commenterid: in.GetCommenterid()}, nil
}
