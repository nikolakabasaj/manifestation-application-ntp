package main

import (
	"controller"
	"fmt"
	"net/http"
	"repository"
	"service"
	"time"
	"model"

	"router"
_ "github.com/lib/pq"
	"database/sql"
)

var (
	manifestationRepository repository.ManifestaionRepository = repository.NewManifestationRepository()
	manifestationService    service.ManifestationService       = service.NewManifestationService(manifestationRepository)
	manifestationController controller.ManifestationController = controller.NewManifestationController(manifestationService)

	cardRepository repository.CardRepository = repository.NewCardRepository()
	cardService    service.CardService       = service.NewCardService(cardRepository)
	cardController controller.CardController = controller.NewCardController(cardService)

	commentRepository repository.CommentRepository = repository.NewCommentRepository()
	commentService    service.CommentService       = service.NewCommentService(commentRepository)
	commentController controller.CommentController = controller.NewCommentController(commentService)

	rateRepository repository.RateRepository = repository.NewRateRepository()
	rateService    service.RateService       = service.NewRateService(rateRepository)
	rateController controller.RateController = controller.NewRateController(rateService)

	httpRouter router.Router = router.NewMuxRouter()
)

func runServer() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Running...")
	})

	httpRouter.GET("/manifestation", manifestationController.FindAll)
	httpRouter.POST("/manifestation", manifestationController.Save)
	httpRouter.POST("/manifestation/search", manifestationController.Search)

	// httpRouter.POST("/card", cardController.GetAll)
	// // httpRouter.GET("/card/getAll2", cardController.GetAll2)
	// httpRouter.POST("/card", cardController.Save)

	httpRouter.POST("/comment", commentController.Save)

	httpRouter.POST("/rate", rateController.Save)

	httpRouter.SERVE(port)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "ntp"
)

func insertManifestation() {
	const layout = "2006-01-02"
	d, _ := time.Parse(layout, "2021-05-05")
	fmt.Println(d)

	var manifestation model.Manifestation
	manifestation.Id = "5"
	manifestation.Name = "Manifestation4"
	manifestation.Country = "Estonia"
	manifestation.City = "Talin"
	manifestation.Date = d

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database when return
	defer db.Close()

	var insertStmt string

	date := manifestation.Date.String()
	insertStmt = `insert into "manifestation"("id", "name", "city", "country", "date") values($1, $2, $3, $4, $5)`
	_, e1 := db.Exec(insertStmt, manifestation.Id, manifestation.Name, manifestation.City,
	manifestation.Country, 	date[0:10])
	
	// insertStmt = `insert into "card"("id", "manifestationId", "price", "date") values($1, $2, $3, $4)`
  // _, e1 := db.Exec(insertStmt, "5", "6", 5500, date[0:10])

	// insertStmt = `insert into "rate"("id", "manifestationId", "mark") values($1, $2, $3)`
	// _, e1 := db.Exec(insertStmt, "5", "4", 2)

	CheckError(e1)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runServer()
	//insertManifestation()
}
