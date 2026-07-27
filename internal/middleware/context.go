package middleware

import (
	"context"

	_context "graphophone.identity/internal/context"

	"google.golang.org/grpc"
)

func ContextPropagationUnaryServerInterceptor(cb _context.ContextBuilder) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		return handler(cb.Build(ctx), req)
	}
}
