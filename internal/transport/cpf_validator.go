package transport

import (
	"context"

	"github.com/andreoav/brazilian-utils-service/pkg/api/cpf"
	"github.com/andreoav/brazilian-utils-service/pkg/validators"
)

type ValidateCPFService struct {
	Validator validators.Validator
}

func (s *ValidateCPFService) ValidateCPF(ctx context.Context, request *cpf.ValidatorCPFRequest) (*cpf.ValidatorCPFResponse, error) {
	return &cpf.ValidatorCPFResponse{
		Result: s.Validator.IsValid(request.GetInput()),
	}, nil
}
