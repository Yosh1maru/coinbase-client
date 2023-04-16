package source

// go:generate mockgen -source=source.go -destination=mocks/source_mock.go -package=mocks

import "context"

type SocketSource interface {
	Read(ctx context.Context) (chan []byte, error)
}
