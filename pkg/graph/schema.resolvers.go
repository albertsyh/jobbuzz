package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/b-open/jobbuzz/pkg/graph/generated"
	"github.com/b-open/jobbuzz/pkg/graph/graphmodel"
	"github.com/jinzhu/copier"
	werrors "github.com/pkg/errors"
)

func (r *mutationResolver) RegisterAccount(ctx context.Context, input graphmodel.NewUser) (*graphmodel.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Jobs(ctx context.Context) ([]*graphmodel.Job, error) {
	jobs, err := r.Service.GetJobs()
	if err != nil {
		return nil, werrors.Wrapf(err, "Error in GetJobs")
	}

	// TODO: fix datetime not working
	var graphqlJobs []*graphmodel.Job
	err = copier.Copy(&graphqlJobs, &jobs)
	if err != nil {
		return nil, werrors.Wrapf(err, "Error copying structs")
	}

	return graphqlJobs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
