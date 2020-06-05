package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vikusku/book-freelancer/internal/graph/generated"
	"github.com/vikusku/book-freelancer/internal/graph/model"
)

func (r *mutationResolver) CreateAppointment(ctx context.Context, input model.NewAppointment) (*model.Appointment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Appointment(ctx context.Context, id string) (*model.Appointment, error) {
	return &model.Appointment{
		ID:          "0100",
		Description: "change tires",
		Time: &model.DateTime{
			Date:  "2020-05-06",
			Start: "12:00",
			End:   "12:30",
		},
		User: &model.User{
			ID:       "31423",
			FistName: "Foo",
			LastName: "Bar",
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
