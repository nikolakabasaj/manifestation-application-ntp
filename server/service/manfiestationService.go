package service

import (
	"errors"
	"model"
	"repository"
	"time"

	"github.com/google/uuid"
)

type ManifestationService interface {
	Save(*model.Manifestation) (*model.Manifestation, error)
	FindAll() ([]model.ManifestationResponse, error)
	Search(searchRequest *model.SearchRequest) ([]model.ManifestationResponse, error)
	Validate(manifestation *model.Manifestation) (error)
	ValidateSearchRequest(searchRequest *model.SearchRequest) (error)
}

type manifestationService struct {}

var (
	manifestationRepository repository.ManifestaionRepository
)

func NewManifestationService(repository repository.ManifestaionRepository) ManifestationService {
	manifestationRepository = repository
	return &manifestationService{}
}

func (*manifestationService) Save(manifestation *model.Manifestation) (*model.Manifestation, error) {
	manifestation.Id = uuid.New().String()
	return manifestationRepository.Save(manifestation)
}

func (*manifestationService) FindAll() ([]model.ManifestationResponse, error) {
	return manifestationRepository.FindAll()
}

func (*manifestationService) Search(searchRequest *model.SearchRequest) ([]model.ManifestationResponse, error) {
	return manifestationRepository.Search(searchRequest)
}

func (*manifestationService) Validate(manifestation *model.Manifestation) error {

	if manifestation == nil {
		err := errors.New("The manifestation is not valid.")
		return err
	}

	if manifestation.Name == "" {
		err := errors.New("The name is not valid.")
		return err
	}

	if manifestation.City == "" {
		err := errors.New("The city is not valid.")
		return err
	}

	if manifestation.Country == "" {
		err := errors.New("The country is not valid.")
		return err
	}

	var zeroTime time.Time
	if manifestation.Date == zeroTime {
		err := errors.New("The date is empty.")
		return err
	}

	return nil
}

func (*manifestationService) ValidateSearchRequest(searchRequest *model.SearchRequest) error {
	if searchRequest == nil {
		err := errors.New("The manifestation is not valid.")
		return err
	}

	const layout = "2006-01-02"
	empty, _ := time.Parse(layout, "0001-01-01")
	if searchRequest.DateFrom == empty {
		d, _ := time.Parse(layout, "1990-01-01")
		searchRequest.DateFrom = d
	}

	if searchRequest.DateTo == empty {
		d, _ := time.Parse(layout, "2100-01-01")
		searchRequest.DateTo = d
	}

	if searchRequest.PriceFrom < 0 {
		err := errors.New("The price is not valid.")
		searchRequest.PriceFrom = 0
		return err
		
	}

	if searchRequest.PriceTo < 0 {
		err := errors.New("The price is not valid.")
		searchRequest.PriceTo = 10000
		return err
	}

	return nil
}