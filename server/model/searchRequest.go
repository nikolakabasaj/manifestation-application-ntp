package model

import "time"

type SearchRequest struct {
	Name      string    `json:"name"`
	Country  	string    `json:"country"`
	City  		string    `json:"city"`
	PriceFrom int       `json:"priceFrom"`
	PriceTo   int       `json:"priceTo"`
	DateFrom  time.Time `json:"dateFrom"`
	DateTo    time.Time `json:"dateTo"`

	/*
		value should be:
		- priceAscending
		- priceDescending
		- latest
		- oldest
	*/
	Sort string `json:"sort"`
}