package database

import (
	model "example.com/petproject/models"
	"github.com/jinzhu/gorm"
)

type Dbclient struct {
	Db *gorm.DB
}

type Dbinterface interface {
	CreateCommentDbInteraction(model.Comment)
}

func (s Dbclient) CreateCommentDbInteraction(model.Comment) {
	var comm model.Comment
	s.Db.Save(&comm)
}
