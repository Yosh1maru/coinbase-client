package internal

import (
	"coinbase-client/internal/entity"
	"coinbase-client/internal/source"
	"coinbase-client/internal/storage"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

type App struct {
	client  source.SocketSource
	storage storage.SocketStorage
	logger  *zap.Logger
}

func NewApp(client source.SocketSource, storage storage.SocketStorage, logger *zap.Logger) *App {
	return &App{
		client,
		storage,
		logger,
	}
}

func (a *App) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	readChan, err := a.client.Read(ctx)
	if err != nil {
		a.logger.Error("error while reading the socket connection:", zap.Error(err))
		return
	}

	for data := range readChan {
		var ticker entity.Ticker
		err := json.Unmarshal(data, &ticker)
		if err != nil {
			a.logger.Error("error while unmarshalling json:", zap.Error(err))
			return
		}

		if ticker.Type != "ticker" {
			a.logger.Info("skipping not ticker data")
			continue
		}

		time.Sleep(time.Second)
		err = a.storage.SaveTicker(ticker)
		if err != nil {
			a.logger.Error("error while SaveTicker to storage:", zap.Error(err))
			continue
		}
	}
}
