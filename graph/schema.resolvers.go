package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ponbac/prediction-backend/graph/generated"
	"github.com/ponbac/prediction-backend/graph/model"
)

// UpsertUser is the resolver for the upsertUser field.
func (r *mutationResolver) UpsertUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpsertUser - upsertUser"))
}

// UpsertPrediction is the resolver for the upsertPrediction field.
func (r *mutationResolver) UpsertPrediction(ctx context.Context, input model.PredictionInput) (*model.Prediction, error) {
	panic(fmt.Errorf("not implemented: UpsertPrediction - upsertPrediction"))
}

// UpsertGame is the resolver for the upsertGame field.
func (r *mutationResolver) UpsertGame(ctx context.Context, input model.GameInput) (*model.Game, error) {
	panic(fmt.Errorf("not implemented: UpsertGame - upsertGame"))
}

// UpsertGroup is the resolver for the upsertGroup field.
func (r *mutationResolver) UpsertGroup(ctx context.Context, input model.GroupInput) (*model.Group, error) {
	panic(fmt.Errorf("not implemented: UpsertGroup - upsertGroup"))
}

// UpsertTeam is the resolver for the upsertTeam field.
func (r *mutationResolver) UpsertTeam(ctx context.Context, input model.TeamInput) (*model.Team, error) {
	panic(fmt.Errorf("not implemented: UpsertTeam - upsertTeam"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Games is the resolver for the games field.
func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	panic(fmt.Errorf("not implemented: Games - games"))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// Teams is the resolver for the teams field.
func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented: Teams - teams"))
}

// Predictions is the resolver for the predictions field.
func (r *queryResolver) Predictions(ctx context.Context) ([]*model.Prediction, error) {
	panic(fmt.Errorf("not implemented: Predictions - predictions"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
