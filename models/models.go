package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	Company     string
	Location    string
	Status      string
	Skills      []Skill
	Bio         string
	Experiences []Experience
	Posts       []Post
	connections []Connected
}
type Skill struct {
	gorm.Model
	Technology string
	UserID     uint
}

type Experience struct {
	gorm.Model
	Position    string
	Company     string
	From        time.Time
	To          time.Time
	Description string
	UserID      uint
}

type Post struct {
	gorm.Model
	Text     string
	Comments []Comment
	Likes    []Like
	UserID   uint
}

type Comment struct {
	gorm.Model
	Text        string
	CommenterId uint
	PostID      uint
}
type Like struct {
	gorm.Model
	LikerId uint
	PostID  uint
}

type Connected struct {
	gorm.Model
	User_1 uint
	User_2 uint
	Status string
}

func StartDB() {

	envErr := godotenv.Load(".env")
	//fmt.Println(envErr)
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	db_user := os.Getenv("db_user")
	db_password := os.Getenv("db_password")
	db_name := os.Getenv("db_name")
	conn1 := "user=" + db_user + " password=" + db_password + " dbname=" + db_name + " sslmode=disable"
	db, err := gorm.Open("postgres", conn1)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Experience{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Like{})
	db.AutoMigrate(&Connected{})
	db.AutoMigrate(&Skill{})

	db.Model(&Experience{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	db.Model(&Skill{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	db.Model(&Comment{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "RESTRICT")
	db.Model(&Like{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "RESTRICT")
	db.Model(&Connected{}).AddForeignKey("user_1", "users(id)", "CASCADE", "RESTRICT")
	db.Model(&Connected{}).AddForeignKey("user_2", "users(id)", "CASCADE", "RESTRICT")

	// user2 := User{
	// 	Name:    "AB",
	// 	Email:   "AB025@gmail.com",
	// 	Company: "BC",
	// 	Status:  "Active",
	// 	Experiences: []Experience{
	// 		{Position: "software dev", Company: "BC"},
	// 	},
	// }

	// user1 := User{
	// 	Name:    "gst",
	// 	Email:   "suryagarimella@beautifulcode.in",
	// 	Company: "BC",
	// 	Status:  "Active",
	// 	Experiences: []Experience{
	// 		{Position: "software dev", Company: "BC"},
	// 	},
	// 	Skills: []Skill{
	// 		{Technology: "cpp"},
	// 	},
	// 	Posts: []Post{
	// 		{
	// 			Text: "hello connections i've joined BC as dev intern",
	// 			Comments: []Comment{
	// 				{Text: "Congrats gst", CommenterId: 2, PostID: 1},
	// 			},
	// 			Likes: []Like{
	// {LikerId: 2, PostID: 1},
	// 			},
	// 		},
	// 	},
	// 	connections: []Connected{
	// 		{
	// 			User_1: 1,
	// 			User_2: 2,
	// 			Status: "pending",
	// 		},
	// 	},
	// }

	// // db.Save(&user1)
	// db.Save(&user2)

}
