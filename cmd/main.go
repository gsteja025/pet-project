package main

import (
	"fmt"
	"log"
	"net"
	"os"

	database "example.com/petproject/database"
	mod "example.com/petproject/models"
	pb "example.com/petproject/protos"
	ser "example.com/petproject/services"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":54321" // choosing port number
)

func main() {

	mod.StartDB()
	fmt.Println("inside main.go")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//db connection
	url := os.Getenv("DATABASE_URL")

	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//create new server

	new_server := grpc.NewServer()
	pb.RegisterLinkedinDatabaseCrudServer(new_server, &ser.Linkedinserver{
		Db: database.Dbclient{Db: db},
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
