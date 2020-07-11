package auth

import "context"

// Auth interface can be implemented by types used for unary or streaming rpc
type Auth interface {
	Interceptor(ctx context.Context) (context.Context, error)
}

// [JWT in gRPC example](https://github.com/tegk/grpcMiddlewareAuthentication/blob/9d5d422c9be5694ab4cfd9142ad53aedaa951dfe/authentication/authentication.go#L74)
// Reading: https://medium.com/@tillknuesting/grpc-http-basic-auth-oauth2-bearer-tokens-f088b5a2314
