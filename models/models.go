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
}
type Skill struct {
	Technology string
	UserID     uint
}

type Experience struct {
	ExperienceID uint `gorm:"AUTO_INCREMENT"`
	Position     string
	Company      string
	From         time.Time
	To           time.Time
	Description  string
	UserID       uint
}

type Post struct {
	gorm.Model
	Text     string
	Comments []Comment
	Like     []Likes
	UserID   uint
}

type Comment struct {
	CommentID   uint `gorm:"AUTO_INCREMENT"`
	Text        string
	CommenterId uint
	PostID      uint
}
type Likes struct {
	LikesID uint `gorm:"AUTO_INCREMENT"`
	LikerId uint
	PostID  uint
}

type Connected struct {
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

	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
	db.DropTableIfExists(&Experience{})
	db.CreateTable(&Experience{})
	db.DropTableIfExists(&Post{})
	db.CreateTable(&Post{})
	db.DropTableIfExists(&Comment{})
	db.CreateTable(&Comment{})
	db.DropTableIfExists(&Likes{})
	db.CreateTable(&Likes{})
	db.DropTableIfExists(&Connected{})
	db.CreateTable(&Connected{})
	db.DropTableIfExists(&Skill{})
	db.CreateTable(&Skill{})

	user2 := User{
		Name:    "AB",
		Email:   "AB025@gmail.com",
		Company: "BC",
		Status:  "Active",
		Experiences: []Experience{
			{Position: "software dev", Company: "BC"},
		},
	}

	user1 := User{
		Name:    "gst",
		Email:   "suryagarimella@beautifulcode.in",
		Company: "BC",
		Status:  "Active",
		Experiences: []Experience{
			{Position: "software dev", Company: "BC"},
		},
		Skills: []Skill{
			{Technology: "cpp"},
		},
		Posts: []Post{
			{Text: "hello connections i've joined BC as dev intern"},
		},
	}
	comment := Comment{
		Text: "Congrats gst", CommenterId: 2, PostID: 1,
	}
	like := Likes{
		LikerId: 2, PostID: 1,
	}
	conn := Connected{
		User_1: 1,
		User_2: 2,
		Status: "pending",
	}

	db.Save(&user1)
	db.Save(&user2)
	db.Save(&conn)
	db.Save(&comment)
	db.Save(&like)

}
