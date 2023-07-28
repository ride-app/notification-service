package interceptors

import (
	"context"
	"errors"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func NewAuthInterceptor(ctx context.Context) (*connect.UnaryInterceptorFunc, error) {
	jwksURI := "https://www.googleapis.com/service_accounts/v1/jwk/securetoken@system.gserviceaccount.com"

	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			log.Fatal("There was an error with the jwt.Keyfunc")
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURI, options)

	if err != nil {
		log.WithError(err).Error("Failed to create JWKS from URI", err)
		return nil, err
	}

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Header().Get("authorization") == "" {
				log.Info("No token provided")
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}

			if req.Header().Get("authorization")[:7] != "Bearer " {
				log.Info("Invalid token format")
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid token format"),
				)
			}
			token, err := jwt.Parse(req.Header().Get("authorization")[7:], jwks.Keyfunc)

			if !token.Valid {
				log.Info("Invalid token")
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid token"),
				)
			}

			if err != nil {
				log.Info("Failed to parse token", err)
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New(err.Error()),
				)
			}

			req.Header().Add("uid", token.Claims.(jwt.MapClaims)["user_id"].(string))

			log.Debug("uid from jwt: ", token.Claims.(jwt.MapClaims)["user_id"].(string))

			return next(ctx, req)
		})
	}

	interceptorFunc := connect.UnaryInterceptorFunc(interceptor)

	return &interceptorFunc, nil
}
