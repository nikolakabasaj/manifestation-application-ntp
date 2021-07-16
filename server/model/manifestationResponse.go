package model

import "time"

type ManifestationResponse struct {
	Id 					string 			`json:"id"`
	Name 				string 			`json:"name"`
	City 				string 			`json:"city"`
	Country 		string	 		`json:"country"`
	Date 				time.Time 	`json:"date"`
	Price 			int16				`json:"price"`
	AverageRate float64 		`json:"averageRate"`
	Rates    		[]Rate    	`json:"rates"`
	Comments 		[]Comment 	`json:"comments"`
}