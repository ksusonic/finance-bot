package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ksusonic/finance-bot/internal/model"
)

type ChatsPostgresStorage struct {
	db *sqlx.DB
}

func NewChatsStorage(db *sqlx.DB) *ChatsPostgresStorage {
	return &ChatsPostgresStorage{db: db}
}

func (s *ChatsPostgresStorage) Create(ctx context.Context, Chat model.Chat) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := s.db.ExecContext(
		ctx,
		`INSERT INTO chats (id, chat_name, telegram_chat_id, chat_type) VALUES ($1, $2, $3, $4)`,
		Chat.ID,
		Chat.Name,
		Chat.ChatID,
		Chat.ChatType,
	); err != nil {
		return err
	}
	return nil
}

func (s *ChatsPostgresStorage) GetByTelegramID(ctx context.Context, telegramChatID int64) (*model.Chat, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var chat model.Chat
	if err := conn.GetContext(
		ctx,
		&chat,
		"SELECT * FROM chats WHERE telegram_chat_id = $1 LIMIT 1",
		telegramChatID,
	); err != nil {
		return nil, err
	}

	return &chat, nil
}
