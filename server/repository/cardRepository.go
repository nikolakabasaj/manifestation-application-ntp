package repository


import (
	"database/sql"
	"fmt"
	"model"
	"time"

_ "github.com/lib/pq"
)

type CardRepository interface {
	Save(card *model.Card) (*model.Card, error)
}

type cardRepository struct {}

func NewCardRepository() CardRepository{
	return &cardRepository{}
}

func (*cardRepository) Save(card *model.Card) (*model.Card, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "card"("id", "manifestationId", "price", "date") values($1, $2, $3, $4)`
	_, e := db.Exec(insertStmt, card.Id, card.ManifestationId, card.Price, card.Date)
	CheckError(e)

	fmt.Println("*** Card was successfully added! ***")
	return card, nil
}

func getCardPriceByManifestationId(searchManifestationId string) int16 {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	query := `
	SELECT "id", "manifestationId", "price", "date"
	FROM "card"
	WHERE "manifestationId" = $1
	`
	rows, err := db.Query(query, searchManifestationId)
	CheckError(err)

	var price int16
	for rows.Next() {
		var id, manifestationId string
		var date time.Time
		err = rows.Scan(&id, &manifestationId, &price, &date)
		CheckError(err)
	}

	return price
}


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
