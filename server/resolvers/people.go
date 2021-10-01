package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/graphql/server"
)

func (r *mutationResolver) CreatePerson(ctx context.Context, input ent.CreatePerson) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePerson(ctx context.Context, id int, input ent.UpdatePerson) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetPersonDeceased(ctx context.Context, id int, cause string) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetPersonAlive(ctx context.Context, id int) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ClaimPerson(ctx context.Context, id int) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SurrenderPerson(ctx context.Context, id int) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ArchivePerson(ctx context.Context, id int) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RestorePerson(ctx context.Context, id int) (*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *personResolver) Age(ctx context.Context, obj *ent.Person) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchPeople(ctx context.Context, where *ent.PersonWhereInput) ([]*ent.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

// Person returns server.PersonResolver implementation.
func (r *Resolver) Person() server.PersonResolver { return &personResolver{r} }

type personResolver struct{ *Resolver }
