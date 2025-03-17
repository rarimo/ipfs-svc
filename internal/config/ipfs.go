package config

import (
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	defaultTimeout  = 30 * time.Second
	defaultFileSize = 10 << 20 // 10MB
)

type IPFSConfig interface {
	IPFS() *IPFSParams
}

type ipfsConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func NewIPFSConfig(getter kv.Getter) IPFSConfig {
	return &ipfsConfig{
		getter: getter,
	}
}

type IPFSParams struct {
	NodeAddress string        `fig:"node_address,required"`
	Timeout     time.Duration `fig:"timeout"`
	MaxFileSize int64         `fig:"max_file_size"`
}

func (c *ipfsConfig) IPFS() *IPFSParams {
	return c.once.Do(func() interface{} {
		var cfg IPFSParams
		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "ipfs")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ipfs config"))
		}

		if cfg.Timeout == 0 {
			cfg.Timeout = defaultTimeout
		}

		if cfg.MaxFileSize == 0 {
			cfg.MaxFileSize = defaultFileSize
		}

		return &cfg
	}).(*IPFSParams)
}
