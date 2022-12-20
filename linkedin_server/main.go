package main

import (
	"log"
	"net"

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

func main() {

	mod.StartDB()
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
	pb.RegisterLinkedinDatabaseCrudServer(new_server, &ser.Linkedinserver{
		Db: db,
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
