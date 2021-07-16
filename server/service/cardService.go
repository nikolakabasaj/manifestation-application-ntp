package service

import (
	"errors"
	"model"
	"repository"
	"time"

	"github.com/google/uuid"
)

type CardService interface {
	Save(*model.Card) (*model.Card, error)
	Validate(card *model.Card) (error)
}

type cardService struct {}

var (
	cardRepository repository.CardRepository
)

func NewCardService(repository repository.CardRepository) CardService {
	cardRepository = repository
	return &cardService{}
}

func (*cardService) Save(card *model.Card) (*model.Card, error) {
	card.Id = uuid.New().String()
	return cardRepository.Save(card)
}

func (*cardService) Validate(card *model.Card) error {

	if card == nil {
		err := errors.New("The card is not valid.")
		return err
	}

	if card.ManifestationId == "" {
		err := errors.New("The manifestation id is  not valid.")
		return err
	}

	if card.Price < 0 {
		err := errors.New("The price is not valid.")
		return err
	}

	var zeroTime time.Time
	if card.Date == zeroTime {
		err := errors.New("The date is empty.")
		return err
	}

	return nil
}