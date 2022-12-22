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

	// As a User I should be able to see  all comments on  my post

	// comments, err2 := client.AllCommentsOfPost(ctx, &pb.Post{Id: 1, Text: "hello connections i've joined BC as dev intern", UserID: 1})
	// checkerror(err2)
	// log.Printf("These are comments of your posts")
	// for _, comm := range comments.Allcomments {
	// 	fmt.Println(comm.GetText())
	// }

	// likes, err3 := client.GetPostLikes(ctx, &pb.Post{Id: 1, Text: "hello connections i've joined BC as dev intern", UserID: 1})
	// checkerror(err3)
	// log.Printf("These are likes of your posts")
	// for _, comm := range likes.Users {
	// 	fmt.Println(comm.GetId())
	// }

	// jobs, err4 := client.ConnectWithOtherUser(ctx, &pb.ConnectionRequest{
	// 	Id1: 2,
	// 	Id2: 1,
	// })
	// checkerror(err4)
	// fmt.Println(jobs)

	// connected_users, err5 := client.GetConnectedUsers(ctx, &pb.User{Id: 1, Name: "gst"})
	// checkerror(err5)
	// log.Printf("These are your connected users")
	// for _, user := range connected_users.Users {
	// 	fmt.Println(user.GetId())
	// }

	// like, err6 := client.LikeOtherPost(ctx, &pb.Post{Id: 1, UserID: 2})
	// checkerror(err6)
	// fmt.Println(like)

	us1, err7 := client.SearchUser(ctx, &pb.SearchRequest{Tech: *pb.SearchRequest_cpp.Enum()})
	checkerror(err7)
	log.Printf("These are your relevant searches")
	fmt.Println(us1)

	// new_comment, err8 := client.CreateComment(ctx, &pb.Comment{Text: "Great news", Commenterid: 1, PostID: 1})
	// checkerror(err8)

	// log.Printf("Comment text: %v", new_comment.GetText())

	//get details of all employees
	// AllEmployees, err := client.GetEmployees(ctx, &pb.EmptyEmployee{})
	// if err != nil {
	// 	log.Printf("error getting employees")
	// }

	// for _, emp := range AllEmployees.Employees {
	// 	fmt.Println(emp.GetEmpName(), emp.GetManagerName(), emp.GetDepartmentId())
	// }

	// // updating manager of employee
	// updated_manager, err := client.UpdateManager(ctx, &pb.Employee{EmpName: "KHK", ManagerName: "Ravi"})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(updated_manager)

	// // deleting an employee
	// delete, err := client.DeleteEmployee(ctx, &pb.Employee{EmpName: "sameer"})
	// if err != nil {
	// 	fmt.Println(delete)
	// 	panic(err.Error())
	// }

}

func checkerror(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
