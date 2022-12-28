package database

import (
	model "example.com/petproject/models"
	"github.com/jinzhu/gorm"
)

type Dbclient struct {
	Db *gorm.DB
}

type Dbinterface interface {
	CreateCommentDbInteraction(model.Comment) (model.Comment, error)
	GetConnectedUsersDbInteraction(model.Connected) ([]model.Connected, error)
	ConnectWithOtherUserDbinteraction1([]model.User) (model.Connected, error)
	ConnectWithOtherUserDbinteraction2([]model.User) error
	ConnectWithOtherUserDbinteraction3([]model.User) error
	CreatepostDbInteraction(model.Post) (model.Post, error)
	GetPostCommentsDbinteraction(model.Post) ([]model.Comment, error)
	GetPostLikesDbinteraction(model.Post) ([]model.Like, error)
	LikeOtherPostDbinteraction(model.Like) error
	SearchUserDbinteraction(model.Skill) ([]model.Skill, error)
}

func (s Dbclient) CreateCommentDbInteraction(comm model.Comment) (model.Comment, error) {
	Db := s.Db.Save(&comm)
	return comm, Db.Error
}

func (s Dbclient) GetConnectedUsersDbInteraction(conn model.Connected) ([]model.Connected, error) {
	//	s.Db.Where("user_1 = ?",conn.User_1).Find(&conn)
	result := []model.Connected{}
	Db := s.Db.Where("user_1 = ?", conn.User_1).Find(&result)
	Db1 := s.Db.Where("user_2 = ?", conn.User_1).Find(&result)
	if Db1.Error != nil && Db.Error == nil {
		Db.Error = Db1.Error
	}
	return result, Db.Error
}
func (s Dbclient) ConnectWithOtherUserDbinteraction1(user []model.User) (model.Connected, error) {
	var conn model.Connected
	Db := s.Db.Where("user_1 = ? and user_2 = ?", user[1].ID, user[0].ID).Find(&conn)
	return conn, Db.Error
}

func (s Dbclient) ConnectWithOtherUserDbinteraction2(user []model.User) error {

	Db := s.Db.Save(&model.Connected{User_1: uint(user[1].ID), User_2: uint(user[0].ID), Status: "Connected"})
	return Db.Error

}

func (s Dbclient) ConnectWithOtherUserDbinteraction3(user []model.User) error {

	Db := s.Db.Save(&model.Connected{User_1: uint(user[0].ID), User_2: uint(user[1].ID), Status: "Pending"})
	return Db.Error
}

func (s Dbclient) CreatepostDbInteraction(post model.Post) (model.Post, error) {

	Db := s.Db.Save(&post)
	return post, Db.Error
}

func (s Dbclient) GetPostCommentsDbinteraction(post model.Post) ([]model.Comment, error) {
	allcommen := []model.Comment{}
	Db := s.Db.Where("post_id = ?", post.ID).Find(&allcommen)
	return allcommen, Db.Error
}

func (s Dbclient) GetPostLikesDbinteraction(post model.Post) ([]model.Like, error) {

	allLikes := []model.Like{}
	Db := s.Db.Where("post_id = ?", post.ID).Find(&allLikes)
	return allLikes, Db.Error
}

func (s Dbclient) LikeOtherPostDbinteraction(likes model.Like) error {
	Db := s.Db.Save(&likes)
	return Db.Error
}

func (s Dbclient) SearchUserDbinteraction(skill model.Skill) ([]model.Skill, error) {

	Allskills := []model.Skill{}

	Db := s.Db.Where(&skill).Find(&Allskills)
	return Allskills, Db.Error
}
