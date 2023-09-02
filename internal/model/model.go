package model

import "time"

type User struct {
	ID         int64
	TelegramID int64
	Username   string
	FirstName  string
	LastName   string
	CreatedAt  time.Time
}

type Transaction struct {
	ID        int64
	Name      string
	Amount    int64
	Date      time.Time
	CreatedAt time.Time
}
