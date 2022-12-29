package database

import (
	"database/sql"
	"fmt"
	"testing"

	model "example.com/petproject/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func SetupDBClient() (*gorm.DB, sqlmock.Sqlmock, error) {
	var (
		Db  *sql.DB
		err error
	)

	Db, mock, _ := sqlmock.New()

	db, err := gorm.Open("postgres", Db)
	return db, mock, err
}

func TestCreatepostDbInteraction(t *testing.T) {

	db, mock, err := SetupDBClient()
	if err != nil {
		t.Fatalf("failed to create mock db client: %v", err)
	}
	defer db.Close()
	mockClient := Dbclient{
		Db: db,
	}
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "posts" (.+)`).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1),
	).WillReturnError(err)
	mock.ExpectCommit()
	post, err := mockClient.CreatepostDbInteraction(model.Post{})
	fmt.Println(post)
	if post.ID != 1 || err != nil {
		t.Errorf("failed to create user")
	}

}
