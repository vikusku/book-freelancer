package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/vikusku/book-freelancer/internal/graph/generated"
	"github.com/vikusku/book-freelancer/internal/graph/graphqlmodel"
	"github.com/vikusku/book-freelancer/internal/orm"
	"github.com/vikusku/book-freelancer/internal/orm/ormmodel"
)

func (r *mutationResolver) CreateAppointment(ctx context.Context, input graphqlmodel.NewAppointment) (*graphqlmodel.Appointment, error) {
	res := r.DbConnection.Create(orm.GraphqlToDb(input))

	if res.Error != nil {
		return &graphqlmodel.Appointment{}, res.Error
	}

	if dbValue, ok := res.Value.(*ormmodel.Appointment); ok {
		return orm.DbToGraphql(dbValue), nil
	} else {
		return &graphqlmodel.Appointment{}, errors.New("INTERNAL_SERVER_ERROR")
	}
}

func (r *queryResolver) Appointment(ctx context.Context, id string) (a *graphqlmodel.Appointment, err error) {
	var parsedId uint64

	if parsedId, err = strconv.ParseUint(id, 10, 32); err != nil {
		log.Println("Failed to parse id")
		return nil, errors.New("INTERNAL_SERVER_ERROR")
	}

	var ap ormmodel.Appointment
	res := r.DbConnection.Set("gorm:auto_preload", true).First(&ap, uint32(parsedId))

	if err := res.Error; err == gorm.ErrRecordNotFound {
		return &graphqlmodel.Appointment{}, fmt.Errorf("appointment for id %s not found", id)
	}

	if dbValue, ok := res.Value.(*ormmodel.Appointment); ok {
		return orm.DbToGraphql(dbValue), nil
	} else {
		return &graphqlmodel.Appointment{}, errors.New("INTERNAL_SERVER_ERROR")
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
