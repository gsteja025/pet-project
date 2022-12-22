package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	mocks "example.com/petproject/mocks"
	models "example.com/petproject/models"
	pb "example.com/petproject/protos"
	"github.com/golang/mock/gomock"
)

var required = []models.Connected{
	{User_1: 1,
		User_2: 2},
}
var required1 = models.Connected{
	User_1: 1,
	User_2: 2,
}

func TestGetConnectedUsers(t *testing.T) {

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbinterface(mockcntrl)
	testProd := Linkedinserver{Db: mockProd}

	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	mockProd.EXPECT().GetConnectedUsersDbInteraction(models.Connected{User_1: 1}).Return(required, nil)

	//rr := httptest.NewRecorder()
	expected := &pb.Users{}
	iwantthis := []*pb.User{}
	iwantthis = append(iwantthis, &pb.User{Id: 2})
	expected = &pb.Users{
		Users: iwantthis,
	}
	ans, err := testProd.GetConnectedUsers(ctx, &pb.User{Id: 1})

	if err != nil {
		fmt.Printf("Thers an error")
	}
	fmt.Println(ans)
	fmt.Println(expected)
	// if !reflect.DeepEqual(ans, expected.Users) {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		ans, expected)
	// }
}

// func TestConnectWithOtherUser(t *testing.T) {

// 	mockcntrl := gomock.NewController(t)
// 	defer mockcntrl.Finish()

// 	mockProd := mocks.NewMockDbinterface(mockcntrl)
// 	testProd := Linkedinserver{Db: mockProd}

// 	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	Userslice := []models.User{}
// 	user1 := models.User{}
// 	user2 := models.User{}
// 	user1.ID = 2
// 	user2.ID = 1
// 	Userslice = append(Userslice, user1)
// 	Userslice = append(Userslice, user2)
// 	mockProd.EXPECT().ConnectWithOtherUserDbinteraction1(Userslice).Return(required1, nil)
// 	mockProd.EXPECT().ConnectWithOtherUserDbinteraction2(Userslice).Return(nil)
// 	mockProd.EXPECT().ConnectWithOtherUserDbinteraction3(Userslice).Return(nil)
// 	//rr := httptest.NewRecorder()
// 	expected := &pb.Emptyresponse{}
// 	// iwantthis := []*pb.User{}
// 	// iwantthis = append(iwantthis, &pb.User{Id: 2})
// 	// expected = append(expected, &pb.Users{
// 	// 	Users: iwantthis,
// 	// })
// 	got, err := testProd.ConnectWithOtherUser(ctx, &pb.ConnectionRequest{Id1: 2, Id2: 1})

// 	if err != nil {
// 		fmt.Printf("Thers an error")
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			got, expected)
// 	}
// }
