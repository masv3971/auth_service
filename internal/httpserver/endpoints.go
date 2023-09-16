package httpserver

import (
	"auth_service/internal/apiv1"
	"context"

	"github.com/gin-gonic/gin"
)

func (s *Service) endpointKeyPair(ctx context.Context, c *gin.Context) (interface{}, error) {
	request := &apiv1.KeyPairRequest{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.KeyPair(ctx, request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointToken(ctx context.Context, c *gin.Context) (interface{}, error) {
	request := &apiv1.TokenRequest{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.Token(ctx, request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointValidate(ctx context.Context, c *gin.Context) (interface{}, error) {
	request := &apiv1.ValidateRequest{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.Validate(ctx, request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointStatus(ctx context.Context, c *gin.Context) (interface{}, error) {
	reply, err := s.apiv1.Status(ctx)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
