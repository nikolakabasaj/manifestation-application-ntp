package service

import (
	"errors"
	"model"
	"repository"

	"github.com/google/uuid"
)

type CommentService interface {
	Save(*model.Comment) (*model.Comment, error)
	Validate(*model.Comment) (error)
}

type commentService struct {}

var (
	commentRepository repository.CommentRepository
)

func NewCommentService(repository repository.CommentRepository) CommentService {
	commentRepository = repository
	return &commentService{}
}

func (*commentService) Save(comment *model.Comment) (*model.Comment, error) {
	comment.Id = uuid.New().String()
	return commentRepository.Save(comment)
}

func (*commentService) Validate(comment *model.Comment) error {

	if comment == nil {
		err := errors.New("The comment is not valid.")
		return err
	}

	if comment.ManifestationId == "" {
		err := errors.New("The manifestation id is not valid.")
		return err
	}

	if comment.Content == "" {
		err := errors.New("The content is not valid.")
		return err
	}

	return nil
}