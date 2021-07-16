package model

import "time"

type Manifestation struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Country string `json:"country"`
	Date time.Time `json:"date"`
}