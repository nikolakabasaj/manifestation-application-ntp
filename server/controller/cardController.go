package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
	"fmt"
)

type CardController interface {

}

var (
	cardService service.CardService
)

type cardController struct {}

func NewCardController(service service.CardService) CardController {
	cardService = service
	return &cardController{}
}

func (*cardController) Save(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var card model.Card

	err := json.NewDecoder(request.Body).Decode(&card)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err.Error()})
		return
	}

	errVal:= cardService.Validate(&card)
	if errVal != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: errVal.Error()})
		fmt.Println(errVal)
		return
	}

	result, err := cardService.Save(&card)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err.Error()})
		return
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(result)
}