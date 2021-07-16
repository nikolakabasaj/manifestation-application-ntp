package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
	"fmt"
)

type CommentController interface {
	Save(response http.ResponseWriter, request *http.Request)
}

var (
	commentService service.CommentService
)

type commentController struct {}

func NewCommentController(service service.CommentService) CommentController {
	commentService = service
	return &commentController{}
}

func (*commentController) Save(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var comment model.Comment

	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err.Error()})
		return
	}
	
	errVal := commentService.Validate(&comment)
	if errVal != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: errVal.Error()})
		fmt.Println(errVal)
		return
	}

	result, err := commentService.Save(&comment)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error saving the comment."})
		return
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(result)
}