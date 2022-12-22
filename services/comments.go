package services

import (
	"context"
	"log"

	mod "example.com/petproject/models"
	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

func (s Linkedinserver) CreateComment(ctx context.Context, in *pb.Comment) (*pb.Comment, error) {
	log.Printf("creating new Comment")
	newcom := mod.Comment{
		Text:        in.GetText(),
		CommenterId: uint(in.GetCommenterid()),
		PostID:      uint(in.GetPostID()),
	}
	s.Db.CreateCommentDbInteraction(newcom)
	return &pb.Comment{Id: uint64(newcom.CommentID), Text: in.GetText(), Commenterid: in.GetCommenterid()}, nil
}
