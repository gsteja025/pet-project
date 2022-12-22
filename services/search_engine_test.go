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

// var required2 = models.Post{Text: "GST", UserID: 1}
var required5 = models.Skill{Technology: "CPP", UserID: 1}

// var required4 = models.Likes{LikerId: 2}

func TestSearchUser(t *testing.T) {

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbinterface(mockcntrl)
	testProd := Linkedinserver{Db: mockProd}

	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// skill := ser.Skill{}
	// skill.Technology = "CPP"
	// skill.UserID = 1
	mockProd.EXPECT().SearchUserDbinteraction(ser.Skill{Technology: "CPP"}).Return([]models.Skill{required5}, nil)

	//rr := httptest.NewRecorder()
	expected := &pb.Users{}
	iwantthis := []*pb.User{}
	iwantthis = append(iwantthis, &pb.User{Id: 1})
	expected.Users = iwantthis
	ans, err := testProd.SearchUser(ctx, &pb.SearchRequest{Tech: *pb.SearchRequest_CPP.Enum()})
	//  var got = a
	// fmt.Println(ans)
	if err != nil {
		fmt.Printf("Thers an error")
	}
	fmt.Println(ans)
	fmt.Println(expected)
	if !reflect.DeepEqual(ans.Users, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			ans, expected)
	}
}
