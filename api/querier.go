package api

import (
	"context"
	ndedb "github.com/ricky8221/NDE_DB/db/sqlc"
)

type Querier interface {
	CreateCompany(ctx context.Context, arg ndedb.CreateCompanyParams) (ndedb.Company, error)
}

var _ Querier = (*ndedb.Queries)(nil)
