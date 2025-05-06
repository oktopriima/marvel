package session

import (
	"context"
	"encoding/json"

	"google.golang.org/api/option"
)

type googlePKG struct {
	ctx context.Context
	cfg *Config
}

func New(ctx context.Context, cfg *Config) Contract {
	return &googlePKG{
		cfg: cfg,
		ctx: ctx,
	}
}

func (g *googlePKG) Context() context.Context {
	return g.ctx
}

func (g *googlePKG) GetConfig() *Config {
	return g.cfg
}

func (g *googlePKG) Option() []option.ClientOption {
	data, _ := json.Marshal(g.cfg)
	return optionCredential(data)
}

func optionCredential(cfg []byte) []option.ClientOption {
	return []option.ClientOption{
		option.WithCredentialsJSON(cfg),
	}
}
