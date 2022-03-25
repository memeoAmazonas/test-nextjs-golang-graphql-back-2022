package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/controller"
	"math/rand"

	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/generated"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input *model.NewBook) (*model.Book, error) {
	book := &model.Book{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Title: input.Title,
		Author: &model.User{
			ID:   input.UserID,
			Name: input.Name,
		},
	}
	controller.Save(book)
	return book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	books := controller.FindAll()
	return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
