//go:generate mockgen -source=$GOFILE -destination=$PROJECT_DIR/internal/generated/mock/mock_$GOPACKAGE/$GOFILE

package postgres

import (
	"database/sql"

	"go.uber.org/dig"
)

type (
	TransactionRepository interface {
		BeginTx() (*sql.Tx, error)
		RollbackTx(tx *sql.Tx) error
		CommitTx(tx *sql.Tx) error
	}

	TransactionRepoImpl struct {
		dig.In
		*sql.DB
	}
)

func NewTransactionRepo(impl TransactionRepoImpl) TransactionRepository {
	return &impl
}

func (r *TransactionRepoImpl) BeginTx() (*sql.Tx, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func (r *TransactionRepoImpl) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}

func (r *TransactionRepoImpl) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}
