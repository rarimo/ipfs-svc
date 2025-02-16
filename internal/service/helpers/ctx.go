package helpers

import (
	"context"
	"net/http"

	"github.com/Dmytro-Hladkykh/ipfs-svc/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	configCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxConfig(cfg config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, cfg)
	}
}

func Config(r *http.Request) config.Config {
	return r.Context().Value(configCtxKey).(config.Config)
}
