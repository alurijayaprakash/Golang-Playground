package comment

import (
	"context"
	"errors"
	"fmt"
)

/*
we are declaring errors here to hide underlayer implementations when we get errors
*/
var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not Implemented")
)

// Comment structure for service layer
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

/*
Store - Inteface of Repository Layer
This interface define all of the methods that our service  needs in order to operate
Simply we are decoupuling the dependency between service layer and repository layer.
so service layer need not to know the implementation of repository layer.
we can simply change db by this way by implementing all methods in this interface
Finally we need to pass this store to Service Struct
*/
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service  - is main struct for the Service layer, logic will build on top of this.
type Service struct {
	Store Store
}

// NewService - returns a pointers to new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - get comment by ID (Service Layer)
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Printf("Retriving a Comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
		/* serice layer got error from repository layer, so if we pass err as it is to transport layer,
		that will revel the underlaying implementation of repository layer
		so here we have declared errors above. so we are passing ErrFetchingComment
		*/

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
