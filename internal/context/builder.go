package context

import (
	"context"

	"graphophone.identity/internal/config"
)

type ContextBuilder interface {
	Build(ctx context.Context) context.Context
}

type contextBuilder struct {
	customContext *customContext
}

func NewContextBuilder(cfg config.Config) (ContextBuilder, error) {
	customContext, err := newCustomContext(cfg)
	if err != nil {
		return nil, err
	}
	return &contextBuilder{
		customContext,
	}, nil
}

func (c *contextBuilder) Build(ctx context.Context) context.Context {
	newCustomContext := *c.customContext
	newCustomContext.Context = ctx
	return newCustomContext
}
