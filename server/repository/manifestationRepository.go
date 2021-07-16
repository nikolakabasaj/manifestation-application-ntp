package repository

import (
	"database/sql"
	"fmt"
	"model"
	"time"
	
_ "github.com/lib/pq"
)


type ManifestaionRepository interface {
	Save(*model.Manifestation) (*model.Manifestation, error)
	FindAll() ([]model.ManifestationResponse, error)
	Search(*model.SearchRequest) ([]model.ManifestationResponse, error)
}

type manifestationRepository struct { }

func NewManifestationRepository() ManifestaionRepository {
	return &manifestationRepository{}
}

func (*manifestationRepository)  Save(manifestation *model.Manifestation) (*model.Manifestation, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "Manifestation"("id", "name", "city", "country", "date") values($1, $2, $3, $4, $5)`
	_, e := db.Exec(insertStmt, manifestation.Id, manifestation.Name, manifestation.City, manifestation.Country, manifestation.Date)
	CheckError(e)

	fmt.Println("*** Mnifestation was successfully added! ***")
	return manifestation, nil
}

func(*manifestationRepository) FindAll() ([]model.ManifestationResponse, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT "id", "name", "city", "country", "date" FROM "manifestation"`)
	CheckError(err)
	defer rows.Close()

	var manifestations []model.ManifestationResponse
	for rows.Next() {
		var id, name, city, country string
		var date time.Time
		var price int16
		var avgRate float64
		
		err = rows.Scan(&id, &name, &city, &country, &date)
		CheckError(err)
	
		price = getCardPriceByManifestationId(id)
		rates:=  getRatesByManifestationId(id)
		avgRate = calculateAverageRate(rates)
		comments:= getCommentsByManifestationId(id)
		manifestations = append(manifestations, model.ManifestationResponse{Id: id, Name: name, City: city, Country: country, Date: date, Price: price, AverageRate: avgRate, Rates: rates, Comments: comments})
	}

	return manifestations, nil
}

func (*manifestationRepository) Search(searchRequest *model.SearchRequest) ([]model.ManifestationResponse, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	searchPreprocess(searchRequest)
	query := getSearchQuery(searchRequest)
	rows, err := db.Query(query, searchRequest.PriceFrom, searchRequest.PriceTo, searchRequest.DateFrom, searchRequest.DateTo)
	CheckError(err)

	defer rows.Close()

	manifestations := []model.ManifestationResponse{}
	for rows.Next() {
		var id, name, city, country, cardId, manifestationId string
		var date time.Time
		var price int16
		var avgRate float64

		err = rows.Scan(&id, &name, &city, &country, &cardId, &manifestationId, &price, &date)
		CheckError(err)

		price = getCardPriceByManifestationId(id)
		rates:=  getRatesByManifestationId(id)
		comments:= getCommentsByManifestationId(id)
		avgRate = calculateAverageRate(rates)
		manifestations = append(manifestations, model.ManifestationResponse{Id: id, Name: name, City: city, Country: country, Date: date, Price: price, AverageRate: avgRate, Rates: rates, Comments: comments})
	}

	manifestations = filterByDate(manifestations, searchRequest.DateFrom, searchRequest.DateTo)
	manifestations = sortByDate(manifestations, searchRequest.Sort)

	return manifestations, nil
}

func sum(array []model.Rate) float64 {  
 result := 0.0  
 for _, e := range array {  
	result +=  float64(e.Mark)  
 }  
 return result  
}  

func calculateAverageRate(rates []model.Rate) float64 {
	var sum = sum(rates)
	if sum == 0.0  {
		return sum
	} 
	return sum/float64(len(rates))
}