package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	mocks "example.com/petproject/mocks"
	models "example.com/petproject/models"
	ser "example.com/petproject/models"
	pb "example.com/petproject/protos"
	"github.com/golang/mock/gomock"
)

// var required = []models.Connected{
// 	{User_1: 1,
// 		User_2: 2},
// }
// var required1 = models.Connected{
// 	User_1: 1,
// 	User_2: 2,
// }

var required2 = models.Post{Text: "GST", UserID: 1}
var required3 = models.Comment{Text: "Congrats gst"}
var required4 = models.Likes{LikerId: 2}

func TestCreatepost(t *testing.T) {

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbinterface(mockcntrl)
	testProd := Linkedinserver{Db: mockProd}

	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	mockProd.EXPECT().CreatepostDbInteraction(models.Post{Text: "GST", UserID: 1}).Return(required2, nil)

	//rr := httptest.NewRecorder()
	expected := &pb.Post{
		Text:   "GST",
		UserID: 1,
	}

	ans, err := testProd.Createpost(ctx, &pb.NewPost{Text: "GST", UserID: 1})
	//  var got = a
	// fmt.Println(ans)
	if err != nil {
		fmt.Printf("Thers an error")
	}
	fmt.Println(ans)
	fmt.Println(expected)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			ans, expected)
	}
}

func TestGetPostComments(t *testing.T) {

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbinterface(mockcntrl)
	testProd := Linkedinserver{Db: mockProd}

	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	pos := ser.Post{}
	pos.ID = 1
	mockProd.EXPECT().GetPostCommentsDbinteraction(pos).Return([]models.Comment{required3}, nil)

	//rr := httptest.NewRecorder()
	expected := &pb.Comments{}
	iwantthis := []*pb.Comment{}
	iwantthis = append(iwantthis, &pb.Comment{Text: "Congrats gst"})
	expected.Allcomments = iwantthis
	ans, err := testProd.GetPostComments(ctx, &pb.Post{Id: 1})
	//  var got = a
	// fmt.Println(ans)
	if err != nil {
		fmt.Printf("Thers an error")
	}
	fmt.Println(ans)
	fmt.Println(expected)
	if !reflect.DeepEqual(ans.Allcomments, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			ans, expected)
	}
}

func TestGetPostLikes(t *testing.T) {

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbinterface(mockcntrl)
	testProd := Linkedinserver{Db: mockProd}

	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	pos := ser.Post{}
	pos.ID = 1
	mockProd.EXPECT().GetPostLikesDbinteraction(pos).Return([]models.Likes{required4}, nil)

	//rr := httptest.NewRecorder()
	expected := &pb.Users{}
	iwantthis := []*pb.User{}
	iwantthis = append(iwantthis, &pb.User{Id: 2})
	expected.Users = iwantthis

	ans, err := testProd.GetPostLikes(ctx, &pb.Post{Id: 1})
	//  var got = a
	// fmt.Println(ans)
	if err != nil {
		fmt.Printf("There an error")
	}
	fmt.Println(ans)
	fmt.Println(expected)
	if !reflect.DeepEqual(ans.Users, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			ans, expected)
	}
}

func TestLikeOtherPost(t *testing.T) {
	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbinterface(mockcntrl)
	testProd := Linkedinserver{Db: mockProd}

	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	pos := ser.Likes{}
	pos.PostID = 1
	pos.LikerId = 2
	mockProd.EXPECT().LikeOtherPostDbinteraction(pos).Return(nil)

	//rr := httptest.NewRecorder()
	expected := &pb.Emptyresponse{}

	ans, err := testProd.LikeOtherPost(ctx, &pb.Request{PostID: 1, LikerID: 2})
	//  var got = a
	// fmt.Println(ans)
	if err != nil {
		fmt.Printf("There an error")
	}
	fmt.Println(ans)
	fmt.Println(expected)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			ans, expected)
	}
}
