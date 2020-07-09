package auth

import (
	"context"
	"encoding/base64"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type basicAuth struct {
	username string
	password string
}

// Interceptor provides basic authentication
func (ba *basicAuth) Interceptor(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "no basic header found: %v", err)
	}
	auth := ba.username + ":" + ba.password
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	if token != authEncoded {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth credentials: %v", err)
	}
	newCtx := context.WithValue(ctx, "Authenticated", true)
	return newCtx, nil
}

// NewBasicAuth creates an instance of Auth
func NewBasicAuth() Auth {
	return &basicAuth{
		"Alice",
		"password123",
	}
}
