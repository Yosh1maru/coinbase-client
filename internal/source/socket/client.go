package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"sync"
)

type Client struct {
	connect              *websocket.Conn
	channels, productIds []string
	sourceType           string
	logger               *zap.Logger
}

func NewClient(connection *websocket.Conn, sourceType string, channels, productIds []string, logger *zap.Logger) *Client {
	return &Client{
		connection,
		channels,
		productIds,
		sourceType,
		logger,
	}
}

func (c *Client) Read(ctx context.Context) (chan []byte, error) {
	var mu sync.Mutex
	channels := make([]chan []byte, 3)
	for i, productId := range c.productIds {
		requestData, err := json.Marshal(map[string]interface{}{
			"type":        c.sourceType,
			"product_ids": []string{productId},
			"channels":    c.channels,
		})
		if err != nil {
			return nil, fmt.Errorf("error while Marshal data in Read: %v", err)
		}

		channels[i] = c.readProductFromConnection(ctx, &mu, requestData)
	}

	return c.mergeChannels(channels...), nil
}

func (c *Client) readProductFromConnection(ctx context.Context, mu *sync.Mutex, readData []byte) chan []byte {
	ch := make(chan []byte)
	mu.Lock()
	_, err := c.connect.Write(readData)
	if err != nil {
		return nil
	}
	mu.Unlock()

	readSocket := func(ch chan []byte) {
		c.logger.Info("reading goroutine starts")
		for {
			select {
			case <-ctx.Done():
				c.logger.Info("reading goroutine ends")
				close(ch)
				return
			default:
				var message = make([]byte, 512)
				n, err := c.connect.Read(message)
				if err != nil {
					continue
				}
				ch <- message[:n]
			}
		}
	}

	go readSocket(ch)

	return ch
}

func (c *Client) mergeChannels(channels ...chan []byte) chan []byte {
	var wg sync.WaitGroup
	out := make(chan []byte)
	c.logger.Info("out goroutine starts")

	output := func(c <-chan []byte) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(channels))
	for _, channel := range channels {
		go output(channel)
	}

	go func() {
		wg.Wait()
		c.logger.Info("out goroutine ends")
		close(out)
	}()

	return out
}
