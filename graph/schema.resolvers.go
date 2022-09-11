package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ponbac/prediction-backend/graph/generated"
	"github.com/ponbac/prediction-backend/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// CreateGame is the resolver for the createGame field.
func (r *mutationResolver) CreateGame(ctx context.Context, input model.CreateGameInput) (*model.Game, error) {
	panic(fmt.Errorf("not implemented: CreateGame - createGame"))
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	panic(fmt.Errorf("not implemented: CreateGroup - createGroup"))
}

// CreateTeam is the resolver for the createTeam field.
func (r *mutationResolver) CreateTeam(ctx context.Context, input model.CreateTeamInput) (*model.Team, error) {
	panic(fmt.Errorf("not implemented: CreateTeam - createTeam"))
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
