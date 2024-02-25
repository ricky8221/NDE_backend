package api

import (
	"context"
	"database/sql"
	"fmt"
	ndedb "github.com/ricky8221/NDE_DB/db/sqlc"
)

// Store provide all functions to execute db queries and transactions
type Store interface {
	Querier
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: ndedb.New(db),
	}
}

// SQLStore provide all functions to execute SQL queries and transactions
type SQLStore struct {
	*ndedb.Queries
	db *sql.DB
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(queries *ndedb.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := ndedb.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
