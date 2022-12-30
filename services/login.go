package services

import (
	"context"
	"log"
	"time"

	pb "example.com/petproject/protos"
	_ "github.com/lib/pq"
)

func (s Linkedinserver) CreateToken(ctx context.Context, in *pb.User) (*pb.Token, error) {
	log.Printf("creating new token called")
	user, err := NewUser(in.Name, in.Email)
	jwtManager := JWTManager{
		SecretKey:     "mysecretkeygst",
		TokenDuration: 5 * time.Minute,
	}
	tokenResponse, err := jwtManager.Generate(user)
	if err != nil {
		log.Println("Error has occured in line 21")
	}

	return &pb.Token{Token: tokenResponse}, nil
}
