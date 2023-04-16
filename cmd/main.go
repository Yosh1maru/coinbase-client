package main

import (
	"coinbase-client/internal"
	"coinbase-client/internal/source/socket"
	"coinbase-client/internal/storage/mysql"
	"context"
	"fmt"
	"github.com/codingconcepts/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	_ "net/http/pprof"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Start application")
	// check active goroutines in the end
	defer func() {
		logger.Info(fmt.Sprintf("count of Goroutines: %v", runtime.NumGoroutine()))
	}()

	var cfg Config
	if err := env.Set(&cfg); err != nil {
		logger.Fatal("Failed to parse config.")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	socketCon, err := socketConnection(cfg.SourceUrl)
	if err != nil {
		logger.Error("error while setup socket connection", zap.Error(err))
		return
	}
	defer socketCon.Close()
	socketClient := socket.NewClient(socketCon, cfg.Type, cfg.Channels, cfg.ProductIds, logger)
	db, err := dbConnection(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName)
	if err != nil {
		logger.Error("error while setup DB connection", zap.Error(err))
		return
	}
	defer db.Close()
	storage := mysql.NewStorage(db)

	application := internal.NewApp(socketClient, storage, logger)
	application.Run(ctx)

	logger.Info("waiting for the completion of subtraction from the channels")
	time.Sleep(time.Second * time.Duration(cfg.SleepDurationTillEnd))
}

func dbConnection(username, password, host, base string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s)/%s", username, password, host, base))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func socketConnection(url string) (*websocket.Conn, error) {
	connect, err := websocket.Dial(url, "", url)
	if err != nil {
		return nil, err
	}

	return connect, nil
}
