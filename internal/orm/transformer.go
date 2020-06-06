package orm

import (
	"github.com/vikusku/book-freelancer/internal/graph/graphqlmodel"
	"github.com/vikusku/book-freelancer/internal/orm/ormmodel"
	"strconv"
)

func GraphqlToDb(input graphqlmodel.NewAppointment) *ormmodel.Appointment{
	return &ormmodel.Appointment{
		Description: input.Description,
		Date:        input.Time.Date,
		Start:       input.Time.Start,
		End:         input.Time.End,
		User:        ormmodel.User{
			FirstName: input.User.FirstName,
			LastName: input.User.LastName,
		},
	}
}

func DbToGraphql(orm *ormmodel.Appointment) *graphqlmodel.Appointment {
	return &graphqlmodel.Appointment{
		ID:          strconv.FormatUint(uint64(orm.ID), 10),
		Description: orm.Description,
		Time: &graphqlmodel.DateTime{
			Date:  orm.Date,
			Start: orm.Start,
			End:   orm.End,
		},
		User:        &graphqlmodel.User{
			ID:       strconv.FormatUint(uint64(orm.User.ID), 10),
			FirstName: orm.User.FirstName,
			LastName: orm.User.LastName,
		},
	}
}