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

func (r *queryResolver) CreateComment(ctx context.Context, input *model.NewComment) (string, error) {
	res := controller.CreateComment(input)
	return res, nil
}

func (r *queryResolver) CreateUser(ctx context.Context, input *model.NewUser) (int, error) {
	res := controller.CreateUser(input)
	return res, nil
}

func (r *queryResolver) CreatePost(ctx context.Context, input *model.NewPost) (int, error) {
	res := controller.CreatePost(input)
	return res, nil
}

func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	res := controller.GetUserByEmail(email)
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
