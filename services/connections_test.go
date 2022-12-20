package services

import (
	"context"
	"testing"
	"time"

	mocks "example.com/petproject/mocks"
	models "example.com/petproject/models"
	pb "example.com/petproject/protos"
	"github.com/golang/mock/gomock"
)

func TestGetConnectedUsers(t *testing.T) {

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbinterface(mockcntrl)
	testProd := Server{Db: mockProd}

	//prod1 := model.Product{Name: "Asus Zenbook 11", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	mockProd.EXPECT().CreateCommentDbInteraction(models.Comment{Text: "this is gst", CommenterId: 2, PostID: 1}).Times(1)
	//rr := httptest.NewRecorder()

	testProd.CreateComment(ctx, &pb.Comment{Text: "this is gst", Commenterid: 2, PostID: 1})

	// Checking status code
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

	// Checking body
	// var got JsonResponse
	// json.NewDecoder(rr.Body).Decode(&got)
	// testProd.CheckErr(err)
	// prod2 := model.Product{Name: "Asus", Description: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate", Quantity: 100, Price: 88000, Image: "lap.jpg"}
	// mockProducts2 := []model.Product{prod2}
	// var mock = JsonResponse{Type: "success", Data: mockProducts2}

	// if !reflect.DeepEqual(got, mock) {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		got, mock)
	// }
}
