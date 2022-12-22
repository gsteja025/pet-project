package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/petproject/protos"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051" // port address for client
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	checkerror(err)
	defer conn.Close()

	client := pb.NewLinkedinDatabaseCrudClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// As a User I should be able to create posts

	new_post, err := client.Createpost(ctx, &pb.NewPost{Text: "we're hiring for software dev role", UserID: 1})
	checkerror(err)

	log.Printf("Post text: %v", new_post.GetText())

	// As a User I should be able to  see other connected users

	us1, err7 := client.SearchUser(ctx, &pb.SearchRequest{Tech: *pb.SearchRequest_CPP.Enum()})
	checkerror(err7)
	log.Printf("These are your relevant searches")
	fmt.Println(us1)

}

func checkerror(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
