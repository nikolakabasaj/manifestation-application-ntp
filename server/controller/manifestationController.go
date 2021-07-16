package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
	"fmt"
)

type ManifestationController interface {
	Save(response http.ResponseWriter, request *http.Request)
	FindAll(response http.ResponseWriter, request *http.Request)
	Search(response http.ResponseWriter, request *http.Request)
}

var (
	manifestationService service.ManifestationService
)

type manifestationController struct {}

func NewManifestationController(service service.ManifestationService) ManifestationController {
	manifestationService = service
	return &manifestationController{}
}

func (*manifestationController) Save(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var manifestation model.Manifestation

	err := json.NewDecoder(request.Body).Decode(&manifestation)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(response).Encode(model.ServiceError{Message: "Error unmarshaling data"})
	// 	return
	// }
	errVal:= manifestationService.Validate(&manifestation)
	if errVal != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: errVal.Error()})
		fmt.Println(errVal)
		return
	}

	result, err := manifestationService.Save(&manifestation)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error saving the manifestation."})
		return
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(result)
}

func (*manifestationController) FindAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	manifestations, err := manifestationService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting manifestations."})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(manifestations)
}

func (*manifestationController) Search(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var searchRequest model.SearchRequest

	err := json.NewDecoder(request.Body).Decode(&searchRequest)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(response).Encode(model.ServiceError{Message: err.Error()})
	// 	return
	// }

	errVal := manifestationService.ValidateSearchRequest(&searchRequest)
	if errVal != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: errVal.Error()})
		fmt.Println(errVal)
		return 
	}

	manifestations, err := manifestationService.Search(&searchRequest)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(model.ServiceError{Message: "Error getting manifestations."})
	}

	response.WriteHeader((http.StatusOK))
	json.NewEncoder(response).Encode(manifestations)
}