package apiv1

import (
	"auth_service/pkg/logger"
	"auth_service/pkg/model"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/masv3971/gosdjwt"
)

// Client holds the public api object
type Client struct {
	cfg           *model.Cfg
	logger        *logger.Log
	sdjwtClient   *gosdjwt.Client
	defaultClaims gosdjwt.Instructions
}

// New creates a new instance of the public api
func New(ctx context.Context, cfg *model.Cfg, logger *logger.Log) (*Client, error) {
	c := &Client{
		cfg:    cfg,
		logger: logger,
		defaultClaims: gosdjwt.Instructions{
			{
				Name:  "iss",
				Value: "https://auth-sdjwt.sunet.se",
			},
			{
				Name:  "exp",
				Value: time.Now().Add(1 * time.Hour).Unix(),
			},
			{
				Name:  "sub",
				Value: uuid.NewString(),
			},
			{
				Name:  "nbf",
				Value: time.Now().Add(-1 * time.Minute).Unix(),
			},
			{
				Name:  "iat",
				Value: time.Now().Unix(),
			},
			{
				Name:  "jti",
				Value: uuid.NewString(),
			},
		},
	}

	var err error
	c.sdjwtClient, err = gosdjwt.New(ctx, gosdjwt.Config{
		JWTType:       "sd-jwt",
		SigningMethod: jwt.SigningMethodES256,
	})
	if err != nil {
		return nil, err
	}

	c.logger.Info("Started")

	return c, nil
}
