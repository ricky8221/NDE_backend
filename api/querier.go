package api

import (
	"context"
	ndedb "github.com/ricky8221/NDE_DB/db/sqlc"
)

type Querier interface {
	GetUser(ctx context.Context, arg string) (ndedb.User, error)
	CreateUser(ctx context.Context, arg ndedb.CreateUserParams) (ndedb.User, error)
	CreateCompany(ctx context.Context, arg ndedb.CreateCompanyParams) (ndedb.Company, error)
	GetCompany(ctx context.Context, arg string) (ndedb.Company, error)
}

var _ Querier = (*ndedb.Queries)(nil)
