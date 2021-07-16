package repository

import (

	"model"
	"sort"
	"time"
)


func getSearchQuery(searchRequest *model.SearchRequest) string {
	query := `
	SELECT "m"."id", "m"."name", "m"."city", "m"."country", "c"."id", 
	"c"."manifestationId", "c"."price", "c"."date"
	FROM "manifestation" "m", "card" "c"
	WHERE "m"."id" = "c"."manifestationId"
	AND "c"."price" BETWEEN $1 AND $2
	AND "m"."date" BETWEEN $3 AND $4
	` + generateNameQuery(searchRequest.Name) + generateCityQuery(searchRequest.City) + generateCountryQuery(searchRequest.Country)	+	generateSortPriceQuery(searchRequest.Sort) + generateSortDateQuery(searchRequest.Sort)

	return query
}

func generateNameQuery(name string) string {
	if name != "" {
		return `AND "m"."name" ilike '%` + name + `%'`
	}
	return ``
}

func generateCityQuery(city string) string {
	if city != "" {
		return `AND "m"."city" ilike '%` + city + `%'`
	}
	return ``
}

func generateCountryQuery(country string) string {
	if country != "" {
		return `AND "m"."country" ilike '%` + country + `%'`
	}
	return ``
}

func generateSortPriceQuery(sort string) string {

	if sort == "priceAscending" {
		return `ORDER BY "c"."price" ASC`
	} else if sort == "priceDescending" {
		return `ORDER BY "c"."price" DESC`
	}

	return ``
}

func generateSortDateQuery(sort string) string {
	if sort == "dateAscending" {
		return `ORDER BY "m"."date" ASC`
	} else if sort == "dateDescending" {
		return `ORDER BY "m"."date" DESC`
	}

	return ``
}

func searchPreprocess(search *model.SearchRequest) {
	if search.PriceFrom <= 0 {
		search.PriceTo = 0
	}
	if search.PriceTo <= 0 {
		search.PriceTo = 9999
	}
}


func filterByDate(manifestations []model.ManifestationResponse, dateFrom time.Time, dateTo time.Time) []model.ManifestationResponse {

	for idx, manifestation := range manifestations {
		if !(manifestation.Date.After(dateFrom) && manifestation.Date.Before(dateTo)) {
			return append(manifestations[0:idx], manifestations[idx+1:]...)
		}
	}

	return manifestations
}

func sortByDate(manifestations []model.ManifestationResponse, sortStr string) []model.ManifestationResponse {
	if sortStr == "latest" {
		sort.Slice(manifestations, func(i, j int) bool {
			return manifestations[i].Date.After(manifestations[j].Date)
		})
	} else if sortStr == "oldest" {
		sort.Slice(manifestations, func(i, j int) bool {
			return manifestations[i].Date.Before(manifestations[j].Date)
		})
	}

	return manifestations

}