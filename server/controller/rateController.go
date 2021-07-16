package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
	"fmt"
)

type RateController interface {
	Save(response http.ResponseWriter, request *http.Request)
}

var (
	rateService service.RateService
)

type rateController struct {}

func NewRateController(service service.RateService) RateController {
	rateService = service
	return &rateController{}
}

func (*rateController) Save(response http.ResponseWriter, request *http.Request) {
	
	response.Header().Set("Content-Type", "application/json")
	var rate model.Rate

	err := json.NewDecoder(request.Body).Decode(&rate)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: err.Error()})
		return
	}

	errVal := rateService.Validate(&rate)
	if errVal != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: errVal.Error()})
		fmt.Println(errVal)
		return
	}

	result, err := rateService.Save(&rate)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error saving the rate."})
		return
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(result)
}