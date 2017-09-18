package middleware

import (
	"context"
	"net/http"
	"strings"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/hashknife/geo-api/endpoints"
)

const (
	// AuthHeader holds the header field name
	AuthHeader = "X-Hashknife-Token"

	// Path
	Path = "Request-Path"
)

// HashknifeRequestAuthenticator
type HashknifeRequestAuthenticator struct {
	authToken    *string
	userAgent    *string
	skippedPaths []string
}

// NewHashknifeRequestAuthenticator
func NewHashknifeRequestAuthenticator(authToken, userAgent *string) *HashknifeRequestAuthenticator {
	return &HashknifeRequestAuthenticator{
		authToken: authToken,
		userAgent: userAgent,
		skippedPaths: []string{
			"/geo-api/healthcheck",
			"/api/healthcheck",
		},
	}
}

// skipAuth is used to not check for an authentication token for a given path
func (a *HashknifeRequestAuthenticator) skipAuth(ctx context.Context) bool {
	if a.authToken == nil || a.userAgent == nil {
		return true
	}
	path, ok := ctx.Value(Path).(string)
	if !ok {
		return false
	}
	for _, s := range a.skippedPaths {
		if strings.HasPrefix(path, s) {
			return true
		}
	}
	return false
}

// verifyTokenHeader concludes that the given headers are or aren't valid
func (a *HashknifeRequestAuthenticator) verifyTokenHeader(ctx context.Context) error {
	if a.skipAuth(ctx) {
		return nil
	}
	authToken, ok := ctx.Value(AuthHeader).(string)
	userAgent, ok := ctx.Value("User-Agent").(string)
	if !ok {
		return endpoints.NewForbiddenError()
	} else if authToken != *a.authToken || userAgent != *a.userAgent {
		return endpoints.NewForbiddenError()
	}
	return nil
}

// EndpointAuthenticate provides an endpoint middleware
func (a *HashknifeRequestAuthenticator) EndpointAuthenticate() kitendpoint.Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (i interface{}, err error) {
			authErr := a.verifyTokenHeader(ctx)
			if authErr != nil {
				return nil, authErr
			}
			return next(ctx, request)
		}
	}
}

// KitServerBefore
func KitServerBefore(ctx context.Context, r *http.Request) context.Context {
	ctx = context.WithValue(ctx, AuthHeader, r.Header.Get(AuthHeader))
	ctx = context.WithValue(ctx, "User-Agent", r.Header.Get("User-Agent"))
	ctx = context.WithValue(ctx, Path, r.URL.Path)
	return ctx
}
