package mysql

import (
	"coinbase-client/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

import (
	"context"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) SaveTicker(ticker entity.Ticker) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	date, err := time.Parse("2006-01-02T15:04:04.000000Z", ticker.Datetime)
	if err != nil {
		return fmt.Errorf("error while Parse ticker datetime: %v", err)
	}
	_, err = s.db.QueryContext(ctx, "INSERT INTO ticks (timestamp, symbol, best_bid, best_ask) VALUES (?, ?, ?, ?)",
		date.Unix(),
		ticker.Symbol,
		ticker.BestBid,
		ticker.BestAsk,
	)
	if err != nil {
		return fmt.Errorf("error while SaveTicker: %v", err)
	}

	return nil
}
