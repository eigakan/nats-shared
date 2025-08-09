package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/eigakan/nats-shared/model"
	"github.com/nats-io/nats.go"
)

type Client struct {
	nc *nats.Conn
}

func NewClient(host string, port int) (*Client, error) {
	nc, err := nats.Connect(fmt.Sprintf("nats://%s:%d", host, port))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to NATS server: %w", err)
	}
	return &Client{
		nc: nc,
	}, nil
}

func (c *Client) Close() error {
	if c.nc == nil {
		return fmt.Errorf("NATS connection is nil")
	}

	c.nc.Close()
	return nil
}

func (c *Client) Drain() error {
	if c.nc == nil {
		return fmt.Errorf("NATS connection is nil")
	}

	return c.nc.Drain()
}

func (c *Client) RespondErr(msg *nats.Msg, errText string) error {
	var resDto model.NatsResponse[any]
	resDto.Err = errText

	res, err := json.Marshal(resDto)
	if err != nil {
		return fmt.Errorf("Failed to marshal error response: %w", err)
	}

	err = msg.Respond(res)

	if err != nil {
		return fmt.Errorf("Failed to send error response via nats: %w", err)
	}

	return nil
}

func (c *Client) Respond(msg *nats.Msg, data any) error {
	var resDto model.NatsResponse[any]
	resDto.Data = data
	resDto.Status = true

	res, err := json.Marshal(resDto)
	if err != nil {
		return fmt.Errorf("Failed to marshal response: %w", err)
	}

	err = msg.Respond(res)

	if err != nil {
		return fmt.Errorf("Failed to send error response via nats: %w", err)
	}

	return nil
}

func (c *Client) Subscribe(topic string, handler nats.MsgHandler) (*nats.Subscription, error) {
	if c.nc == nil {
		return nil, fmt.Errorf("NATS connection is nil")
	}

	sub, err := c.nc.Subscribe(topic, handler)
	if err != nil {
		return nil, fmt.Errorf("Failed to subscribe to subject %s: %w", topic, err)
	}

	return sub, nil
}

func (c *Client) Request(topic string, data any, timeout time.Duration) (*nats.Msg, error) {
	if c.nc == nil {
		return nil, fmt.Errorf("NATS connection is nil")
	}

	req, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal request data: %w", err)
	}

	msg, err := c.nc.Request(topic, req, timeout)
	if err != nil {
		return nil, fmt.Errorf("Failed to send request to subject %s: %w", topic, err)
	}

	return msg, nil
}
