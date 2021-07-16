package repository

import (
	"database/sql"
	"fmt"
	"model"

_ "github.com/lib/pq"
)

type RateRepository interface {
	Save(*model.Rate) (*model.Rate, error)
}

type rateRepository struct {}

func NewRateRepository() RateRepository {
	return &rateRepository{}
}

func (*rateRepository) Save(rate *model.Rate) (*model.Rate, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "rate"("id", "manifestationId", "mark") values($1, $2, $3)`
	_, e := db.Exec(insertStmt, rate.Id, rate.ManifestationId, rate.Mark)
	CheckError(e)

	fmt.Println("*** Rate was successfully added! ***")
	return rate, nil
}


func getRatesByManifestationId(id string) []model.Rate {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	query := `
	SELECT "id", "manifestationId", "mark"
	FROM "rate"
	WHERE "manifestationId" = $1
	`
	rows, err := db.Query(query, id)
	CheckError(err)

	defer rows.Close()

	var rates = []model.Rate{}

	for rows.Next() {
		var id, manifestationId string
		var mark int

		err = rows.Scan(&id, &manifestationId, &mark)
		CheckError(err)

		rates = append(rates, model.Rate{Id: id, ManifestationId: manifestationId, Mark: mark})
	}

	return rates
}