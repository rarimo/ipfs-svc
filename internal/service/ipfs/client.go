package ipfs

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Client struct {
	sh      *shell.Shell
	timeout time.Duration
}

type ClientOpts struct {
	NodeAddress string
	Timeout     time.Duration
}

func NewClient(opts ClientOpts) *Client {
	sh := shell.NewShell(opts.NodeAddress)

	// set timeout for client
	sh.SetTimeout(opts.Timeout)

	return &Client{
		sh:      sh,
		timeout: opts.Timeout,
	}
}

func (c *Client) UploadJSON(ctx context.Context, data json.RawMessage) (string, error) {
	// create context with timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	// convert json to reader
	reader := strings.NewReader(string(data))

	// create channel for result
	type result struct {
		hash string
		err  error
	}
	done := make(chan result, 1)

	go func() {
		hash, err := c.sh.Add(reader, shell.Pin(true))
		done <- result{hash: hash, err: err}
	}()

	// wait for result or timeout
	select {
	case <-ctx.Done():
		return "", errors.Wrap(ctx.Err(), "request timed out")
	case res := <-done:
		if res.err != nil {
			return "", errors.Wrap(res.err, "failed to upload to ipfs")
		}
		return res.hash, nil
	}
}

func (c *Client) UploadFile(ctx context.Context, file io.Reader) (string, error) {
	// create context with timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	// create channel for result
	type result struct {
		hash string
		err  error
	}
	done := make(chan result, 1)

	go func() {
		hash, err := c.sh.Add(file, shell.Pin(true))
		done <- result{hash: hash, err: err}
	}()

	// wait for result or timeout
	select {
	case <-ctx.Done():
		return "", errors.Wrap(ctx.Err(), "request timed out")
	case res := <-done:
		if res.err != nil {
			return "", errors.Wrap(res.err, "failed to upload to ipfs")
		}
		return res.hash, nil
	}
}
