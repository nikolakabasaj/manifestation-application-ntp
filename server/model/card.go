package model

import "time"

type Card struct {
	Id string `json:"id"`
	ManifestationId string `json:"manifestationId"`
	Price int16 `json:"price"`
	Date time.Time `date:"date"`
}