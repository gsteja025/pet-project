package main

import (
	"context"
	"io"
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

	// new_post, err := client.Createpost(ctx, &pb.NewPost{Text: "we're hiring for software dev role", UserID: 1})
	// checkerror(err)

	// log.Printf("Post text: %v", new_post.GetText())

	// As a User I should be able to  see other connected users

	// As a User I should be able to see  all comments on  my post

	comments, err2 := client.GetPostComments(ctx, &pb.PostRequest{Id: 1})
	checkerror(err2)
	log.Printf("These are comments of your posts")
	done := make(chan bool)
	go func() {
		for {
			ele, err := comments.Recv()

			if err == io.EOF {
				done <- true
				break
			}
			if err != nil {
				checkerror(err2)
			}
			log.Println(ele)
		}

	}()
	<-done
	log.Printf("finished")

	// likes, err3 := client.GetPostLikes(ctx, &pb.PostRequest{Id: 1})
	// checkerror(err3)
	// log.Printf("These are likes of your posts")
	// for _, comm := range likes.Users {
	// 	fmt.Println(comm.GetId())
	// }

	// jobs, err4 := client.ConnectWithOtherUser(ctx, &pb.ConnectionRequest{
	// 	Id1: 1,
	// 	Id2: 4,
	// })
	// checkerror(err4)
	// fmt.Println(jobs)

	// token, err := client.CreateToken(ctx, &pb.User{Name: "gst", Email: "gsteja025@gmail.com"})
	// checkerror(err)
	// fmt.Println(token)
	// connected_users, err5 := client.GetConnectedUsers(ctx, &pb.User{Id: 1, Name: "gst"})
	// checkerror(err5)
	// log.Printf("These are your connected users")
	// for _, user := range connected_users.Users {
	// 	fmt.Println(user.GetId())
	// }

	// like, err6 := client.LikeOtherPost(ctx, &pb.Request{PostID: 1, LikerID: 2})
	// checkerror(err6)
	// fmt.Println(like)

	// us1, err7 := client.SearchUser(ctx, &pb.SearchRequest{Tech: *pb.SearchRequest_CPP.Enum()})
	// checkerror(err7)
	// log.Printf("These are your relevant searches")
	// fmt.Println(us1)

	// new_comment, err8 := client.CreateComment(ctx, &pb.NewComment{Text: "Great news", Commenterid: 1, PostID: 1})
	// checkerror(err8)

	// log.Printf("Comment text: %v", new_comment.GetText())

	//get details of all employees

}

func checkerror(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
