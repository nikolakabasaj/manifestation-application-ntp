package service

import (
	"errors"
	"model"
	"repository"

	"github.com/google/uuid"
)

type RateService interface {
	Save(*model.Rate) (*model.Rate, error)
	Validate(rate *model.Rate) (error)
}

type rateService struct {}

var (
	rateRepository repository.RateRepository
)

func NewRateService(repository repository.RateRepository) RateService {
	rateRepository = repository
	return &rateService{}
}

func (*rateService) Save(rate *model.Rate) (*model.Rate, error) {
	rate.Id = uuid.New().String()
	return rateRepository.Save(rate)
}

func (*rateService) Validate(rate *model.Rate) error {
	if rate == nil {
		err := errors.New("The rate is not valid.")
		return err
	}

	if rate.ManifestationId == "" {
		err := errors.New("The manifestation id is not valid.")
		return err
	}
	
	if rate.Mark < 1 || rate.Mark > 5 {
		err := errors.New("The mark must be between 1 and 5.")
		return err
	}

	return nil
}