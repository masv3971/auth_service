package apiv1

import (
	"auth_service/pkg/helpers"
	"auth_service/pkg/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/masv3971/gosdjwt"
)

// KeyPairRequest is the request for sign pdf
type KeyPairRequest struct {
	Algorithm string `json:"algorithm" validate:"required"`
}

// KeyPairReply is the reply for sign pdf
type KeyPairReply struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}

// KeyPair is the request to get a key pair
func (c *Client) KeyPair(ctx context.Context, req *KeyPairRequest) (*KeyPairReply, error) {
	if err := helpers.Check(req, c.logger); err != nil {
		return nil, err
	}

	reply := &KeyPairReply{}

	return reply, nil
}

// TokenRequest is the request for
type TokenRequest struct {
	JWTType            string               `json:"jwt_type" validate:"required"`
	Instructions       gosdjwt.Instructions `json:"instructions" validate:"required"`
	SigningKey         string               `json:"signing_key" validate:"required"`
	PresentationFormat string               `json:"presentation_format"`
}

// TokenReply is the reply for verify pdf
type TokenReply struct {
	JWT string `json:"jwt"`
}

// Token is the handler for verify pdf
func (c *Client) Token(ctx context.Context, req *TokenRequest) (*TokenReply, error) {
	claims := gosdjwt.CombineInstructionsSets(c.defaultClaims, req.Instructions)
	got, err := json.Marshal(claims)
	fmt.Println(string(got))
	signedSDJWT, err := c.sdjwtClient.SDJWT(claims, req.SigningKey)
	fmt.Println(signedSDJWT)
	if err != nil {
		return nil, err
	}
	reply := &TokenReply{
		JWT: signedSDJWT,
	}

	return reply, nil
}

// ValidateRequest is the request for get signed pdf
type ValidateRequest struct {
	Key string `json:"key" validate:"required"`
	JWT string `json:"jwt" validate:"required"`
}

// ValidateReply is the reply for the signed pdf
type ValidateReply struct {
	JWT        jwt.MapClaims `json:"jwt"`
	Validation *gosdjwt.Validation
}

// Validate is the request to get signed pdfs
func (c *Client) Validate(ctx context.Context, req *ValidateRequest) (*ValidateReply, error) {
	jwt, validation, err := gosdjwt.Verify(req.JWT, req.Key)
	if err != nil {
		return nil, err
	}
	reply := &ValidateReply{
		JWT:        jwt,
		Validation: validation,
	}
	return reply, nil
}

// Status return status
func (c *Client) Status(ctx context.Context) (*model.Health, error) {
	probes := model.Probes{}
	//probes = append(probes, c.kv.Status(ctx))

	status := probes.Check("auth_service")

	return status, nil
}
