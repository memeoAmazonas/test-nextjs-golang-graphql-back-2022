package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/generated"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/internal/controller"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	res := controller.GetPost()
	return res, nil
}

func (r *queryResolver) CommentByPost(ctx context.Context, id string) ([]*model.Comment, error) {
	res := controller.GetCommentByPost(id)
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
