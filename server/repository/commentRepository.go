package repository

import (
"database/sql"
"fmt"
"model"

_ "github.com/lib/pq"
)

type CommentRepository interface {
	Save(*model.Comment) (*model.Comment, error)
}

type commentRepository struct {}

func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

func (*commentRepository) Save(comment *model.Comment) (*model.Comment, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "comment"("id", "content", "manifestationId") values($1, $2, $3)`
	_, e := db.Exec(insertStmt, comment.Id, comment.Content, comment.ManifestationId)
	CheckError(e)

	fmt.Println("*** Comment was successfully added! ***")
	return comment, nil
}

func getCommentsByManifestationId(id string) []model.Comment {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	query := `
	SELECT "id", "content", "manifestationId"
	FROM "comment"
	WHERE "manifestationId" = $1
	`
	rows, err := db.Query(query, id)
	CheckError(err)

	defer rows.Close()

	var comments = []model.Comment{}
	for rows.Next() {
		var id, content, manifestationId string

		err = rows.Scan(&id, &content, &manifestationId)
		CheckError(err)

		comments = append(comments, model.Comment{Id: id, Content: content, ManifestationId: manifestationId})
	}

	return comments
}
