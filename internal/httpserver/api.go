package httpserver

import (
	"auth_service/internal/apiv1"
	"auth_service/pkg/model"
	"context"
)

// Apiv1 interface
type Apiv1 interface {
	KeyPair(ctx context.Context, req *apiv1.KeyPairRequest) (*apiv1.KeyPairReply, error)
	Token(ctx context.Context, req *apiv1.TokenRequest) (*apiv1.TokenReply, error)
	Validate(ctx context.Context, req *apiv1.ValidateRequest) (*apiv1.ValidateReply, error)

	Status(ctx context.Context) (*model.Health, error)
}
