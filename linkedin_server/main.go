package main

import (
	"context"
	"log"
	"net"
	"time"

	database "example.com/petproject/database"
	mod "example.com/petproject/models"
	pb "example.com/petproject/protos"
	ser "example.com/petproject/services"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":50051" // choosing port number
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func createUser(userStore ser.UserStore, username, password, role string) error {
	user, err := ser.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore ser.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}
func accessibleRoles() map[string][]string {
	const laptopServicePath = "/techschool.pcbook.LaptopService/"

	return map[string][]string{
		laptopServicePath + "Createpost": {"admin"},
	}
}

func main() {

	mod.StartDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}
	userStore := ser.NewInMemoryUserStore()
	jwtManager := ser.NewJWTManager(secretKey, tokenDuration)

	authServer := ser.NewAuthServer(userStore, jwtManager)

	err1 := seedUsers(userStore)

	if err1 != nil {
		log.Fatal("cannot seed users: ", err)
	}

	//db connection
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	interceptor := ser.NewAuthInterceptor(jwtManager, accessibleRoles())

	//create new server

	new_server := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	pb.RegisterAuthServiceServer(new_server, authServer)

	pb.RegisterLinkedinDatabaseCrudServer(new_server, &ser.Linkedinserver{
		Db: database.Dbclient{Db: db},
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
