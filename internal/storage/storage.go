package storage

// go:generate mockgen -source=storage.go -destination=mocks/storage_mock.go -package=mocks

import (
	"coinbase-client/internal/entity"
)

type SocketStorage interface {
	SaveTicker(ticker entity.Ticker) error
}
