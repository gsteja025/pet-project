package main

import (
	"context"
	"log"
	"net"
	"reflect"

	Email "example.com/petproject/Email"
	model "example.com/petproject/linkedin_models"
	pb "example.com/petproject/protos"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":50051" // choosing port number
)

type linkedinserver struct {
	pb.UnimplementedLinkedinDatabaseCrudServer
	Db *gorm.DB
}

// function to add new employee on server
func (s *linkedinserver) Createpost(ctx context.Context, in *pb.NewPost) (*pb.Post, error) {
	log.Printf("creating new post called")
	newpos := model.Post{
		Text:   in.GetText(),
		UserID: 1,
	}
	s.Db.Save(&newpos)
	return &pb.Post{Text: in.GetText(), Id: uint64(newpos.ID)}, nil
}

func (s *linkedinserver) ConnectedUsers(ctx context.Context, in *pb.User) (*pb.Users, error) {
	log.Printf("Connected users")
	connects := []model.Connected{}
	connects1 := []model.Connected{}
	Finalconnects := []*pb.User{}
	// str1 := "Connected"
	s.Db.Where(&model.Connected{
		User_1: uint(in.GetId()),
		Status: "Connected",
	}).Find(&connects)
	// s.Db.Where("user_2 = ? AND status = ?", in.GetId(), str1).Find(&connects1)
	for _, conn := range connects {
		Finalconnects = append(Finalconnects, &pb.User{Id: uint64(conn.User_2)})
	}

	for _, conn := range connects1 {
		Finalconnects = append(Finalconnects, &pb.User{Id: uint64(conn.User_1)})
	}
	return &pb.Users{Users: Finalconnects}, nil

}

func (s *linkedinserver) AllCommentsOfPost(ctx context.Context, in *pb.Post) (*pb.Comments, error) {
	log.Printf("Getting comments of post")
	allcommen := []model.Comment{}
	Finalcomments := []*pb.Comment{}
	s.Db.Where("post_id = ?", in.GetId()).Find(&allcommen)
	for _, conn := range allcommen {
		Finalcomments = append(Finalcomments, &pb.Comment{Id: uint64(conn.CommentID), Text: conn.Text, Commenterid: uint64(conn.CommenterId)})
	}
	return &pb.Comments{Allcomments: Finalcomments}, nil
}

func (s *linkedinserver) AllUsersWhoLikedspecificPost(ctx context.Context, in *pb.Post) (*pb.Users, error) {
	log.Printf("Getting likes of post")
	allLikes := []model.Likes{}
	FinalLikes := []*pb.User{}
	s.Db.Where("post_id = ?", in.GetId()).Find(&allLikes)
	for _, conn := range allLikes {
		FinalLikes = append(FinalLikes, &pb.User{Id: uint64(conn.LikerId)})
	}
	return &pb.Users{Users: FinalLikes}, nil
}

func (s *linkedinserver) ConnectWithOtherUser(ctx context.Context, in *pb.TwoUsers) (*pb.Emptyresponse, error) {
	log.Printf("YOU NOW HAVE A NEW CONNECTION")
	// var allconnectsids []uint64
	// cnt := 0
	str := "pending"
	var conn model.Connected
	conn1 := conn
	usermodel := model.User{}
	s.Db.Where("user_1 = ? and user_2 = ?", in.GetId2(), in.GetId1()).Find(&conn)
	if !reflect.DeepEqual(conn, conn1) && conn.Status == "pending" {
		str = "Connected"
		s.Db.Save(&model.Connected{User_1: uint(in.GetId2()), User_2: uint(in.GetId1()), Status: str})
		s.Db.Where("id = ?", in.GetId2()).Find(&usermodel)
		Email.SendEmail(usermodel.Email)

	} else {
		s.Db.Save(&model.Connected{User_1: uint(in.GetId1()), User_2: uint(in.GetId2()), Status: str})
	}
	return &pb.Emptyresponse{}, nil
}

func (s *linkedinserver) LikeOtherPosts(ctx context.Context, in *pb.Post) (*pb.Emptyresponse, error) {
	posts := model.Likes{
		PostID:  uint(in.GetId()),
		LikerId: uint(in.UserID),
	}
	s.Db.Save(&posts)
	return &pb.Emptyresponse{}, nil
}

func (S *linkedinserver) SearchForRequiredUserBasedOnTechStack(ctx context.Context, in *pb.Technology) (*pb.Users, error) {
	s1 := model.Skill{
		Technology: in.GetTech(),
	}

	Allskills := []model.Skill{}
	Allusers := []*pb.User{}

	S.Db.Where(&s1).Find(&Allskills)
	for _, ele := range Allskills {
		Allusers = append(Allusers, &pb.User{Id: uint64(ele.UserID)})
	}
	return &pb.Users{Users: Allusers}, nil
}

// function to read employee detail on server
// func (s *employeeServer) GetEmployees(ctx context.Context, in *pb.EmptyEmployee) (*pb.Employees, error) {
// 	log.Printf("Getting employees called")
// 	Employees := []model.Employee{}
// 	FinalEmployees := []*pb.Employee{}
// 	s.Db.Find(&Employees)
// 	for _, emp := range Employees {
// 		FinalEmployees = append(FinalEmployees, &pb.Employee{EmpName: emp.EmpName, DepartmentId: uint64(emp.DepartmentID), ManagerName: emp.ManagerName, Id: uint64(emp.ID)})
// 	}
// 	return &pb.Employees{Employees: FinalEmployees}, nil
// }

// function to update manager on server
// func (s *employeeServer) UpdateManager(ctx context.Context, in *pb.Employee) (*pb.Employee, error) {
// 	log.Printf("update manager called")
// 	s.Db.Model(&model.Employee{}).Where("emp_name=?", in.EmpName).Update("manager_name", in.ManagerName)
// 	return &pb.Employee{EmpName: in.GetEmpName(), ManagerName: in.GetManagerName()}, nil
// }

// // function to delete employee on server
// func (s *employeeServer) DeleteEmployee(ctx context.Context, in *pb.Employee) (*pb.EmptyEmployee, error) {
// 	log.Printf("delete employee called")
// 	s.Db.Where(&model.Employee{EmpName: in.GetEmpName()}).Delete(&model.Employee{})
// 	return &pb.EmptyEmployee{}, nil
// }

func main() {

	model.StartDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//db connection
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//create new server
	new_server := grpc.NewServer()
	pb.RegisterLinkedinDatabaseCrudServer(new_server, &linkedinserver{
		Db: db,
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
