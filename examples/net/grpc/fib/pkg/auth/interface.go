package auth

import "context"

// Auth interface can be implemented by types used for unary or streaming rpc
type Auth interface {
	Interceptor(ctx context.Context) (context.Context, error)
}
