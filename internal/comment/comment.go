package comment

import (
	"context"
	"fmt"
)

// Comment - a representation of the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the
// methods that our service needs
// in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is a the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
// service
// NOTE: Accept interfaces, Return Structs
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")

	// reaches to the repository layer to retrieve a comment
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
