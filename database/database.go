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
	GetConnectedUsersDbInteraction(model.Connected) ([]model.Connected, error)
	ConnectWithOtherUserDbinteraction1([]model.User) (model.Connected, error)
	ConnectWithOtherUserDbinteraction2([]model.User) error
	ConnectWithOtherUserDbinteraction3([]model.User) error
	CreatepostDbInteraction(model.Post) (model.Post, error)
	GetPostCommentsDbinteraction(model.Post) ([]model.Comment, error)
	GetPostLikesDbinteraction(model.Post) ([]model.Likes, error)
	LikeOtherPostDbinteraction(model.Likes) error
	SearchUserDbinteraction(model.Skill) ([]model.Skill, error)
}

func (s Dbclient) CreateCommentDbInteraction(comm model.Comment) {
	s.Db.Save(&comm)
}

func (s Dbclient) GetConnectedUsersDbInteraction(conn model.Connected) ([]model.Connected, error) {
	//	s.Db.Where("user_1 = ?",conn.User_1).Find(&conn)
	result := []model.Connected{}
	s.Db.Where("user_1 = ?", conn.User_1).Find(&result)
	return result, nil
}
func (s Dbclient) ConnectWithOtherUserDbinteraction1(user []model.User) (model.Connected, error) {
	var conn model.Connected
	s.Db.Where("user_1 = ? and user_2 = ?", user[1].ID, user[0].ID).Find(&conn)
	return conn, nil
}

func (s Dbclient) ConnectWithOtherUserDbinteraction2(user []model.User) error {

	s.Db.Save(&model.Connected{User_1: uint(user[1].ID), User_2: uint(user[0].ID), Status: "Connected"})
	return nil

}

func (s Dbclient) ConnectWithOtherUserDbinteraction3(user []model.User) error {

	s.Db.Save(&model.Connected{User_1: uint(user[0].ID), User_2: uint(user[1].ID), Status: "Pending"})
	return nil
}

func (s Dbclient) CreatepostDbInteraction(post model.Post) (model.Post, error) {

	s.Db.Save(&post)
	return post, nil
}

func (s Dbclient) GetPostCommentsDbinteraction(post model.Post) ([]model.Comment, error) {
	allcommen := []model.Comment{}
	s.Db.Where("post_id = ?", post.ID).Find(&allcommen)
	return allcommen, nil
}

func (s Dbclient) GetPostLikesDbinteraction(post model.Post) ([]model.Likes, error) {

	allLikes := []model.Likes{}
	s.Db.Where("post_id = ?", post.ID).Find(&allLikes)
	return allLikes, nil
}

func (s Dbclient) LikeOtherPostDbinteraction(likes model.Likes) error {
	s.Db.Save(&likes)
	return nil
}

func (s Dbclient) SearchUserDbinteraction(skill model.Skill) ([]model.Skill, error) {

	Allskills := []model.Skill{}

	s.Db.Where(&skill).Find(&Allskills)
	return Allskills, nil
}
