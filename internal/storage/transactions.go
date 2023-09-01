package storage

import (
	"context"
	"time"

	"github.com/ksusonic/finance-bot/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type TransactionPostgresStorage struct {
	db *sqlx.DB
}

func NewTransactionsStorage(db *sqlx.DB) *TransactionPostgresStorage {
	return &TransactionPostgresStorage{db: db}
}

func (s *TransactionPostgresStorage) Transactions(ctx context.Context) ([]model.Transaction, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var transactions []dbTransaction
	if err := conn.SelectContext(ctx, &transactions, "SELECT * FROM transactions"); err != nil {
		return nil, err
	}

	return lo.Map(transactions, func(transaction dbTransaction, _ int) model.Transaction { return model.Transaction(transaction) }), nil
}

func (s *TransactionPostgresStorage) TransactionById(ctx context.Context, id int64) (*model.Transaction, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var transaction dbTransaction
	if err := conn.GetContext(ctx, &transaction, "SELECT * FROM transactions WHERE id = $1 LIMIT 1", id); err != nil {
		return nil, err
	}

	return (*model.Transaction)(&transaction), nil
}

func (s *TransactionPostgresStorage) Add(ctx context.Context, transaction *model.Transaction) (int64, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, nil
	}
	defer conn.Close()

	var id int64
	row := conn.QueryRowxContext(
		ctx,
		`INSERT INTO transactions (name, amount, transaction_date) VALUES ($1, $2, $3) RETURNING id`,
		transaction.Name,
		transaction.Amount,
		transaction.Date,
	)
	if err := row.Err(); err != nil {
		return 0, err
	}
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TransactionPostgresStorage) Delete(ctx context.Context, id int64) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, "DELETE FROM transactions WHERE id = $1", id)
	return err
}

type dbTransaction struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Amount    int64     `db:"amount"`
	Date      time.Time `db:"date"`
	CreatedAt time.Time `db:"created_at"`
}
