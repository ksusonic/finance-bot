package storage

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ksusonic/finance-bot/internal/model"
)

type UsersPosgresStorage struct {
	db *sqlx.DB
}

func NewUsersStorage(db *sqlx.DB) *UsersPosgresStorage {
	return &UsersPosgresStorage{db: db}
}

func (s *UsersPosgresStorage) AddUser(ctx context.Context, user model.User) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err = conn.ExecContext(
		ctx,
		`INSERT INTO users (id, telegram_id, username, first_name, last_name) 
				VALUES ($1, $2, $3, $4, $5)
				ON CONFLICT DO NOTHING`,
		user.ID,
		user.TelegramID,
		user.Username,
		user.FirstName,
		user.LastName,
	); err != nil {
		return err
	}
	return nil
}

func (s *UsersPosgresStorage) GetUserByTelegramID(ctx context.Context, telegramId int64) (*model.User, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var user dbUser
	if err := conn.SelectContext(
		ctx,
		&user,
		"SELECT * FROM users WHERE telegram_id = $1 LIMIT 1",
		telegramId,
	); err != nil {
		return nil, err
	}

	return (*model.User)(&user), nil
}

type dbUser struct {
	ID         int64     `db:"id"`
	TelegramID int64     `db:"telegram_id"`
	Username   string    `db:"username"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	CreatedAt  time.Time `db:"created_at"`
}
