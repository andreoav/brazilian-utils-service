package transport

import (
	"context"

	"github.com/andreoav/brazilian-utils-service/pkg/api/cpf"
	"github.com/andreoav/brazilian-utils-service/pkg/generators"
)

type GenerateCPFService struct {
	Generator generators.Generator
}

func (s *GenerateCPFService) GenerateCPF(ctx context.Context, request *cpf.GenerateCPFRequest) (*cpf.GenerateCPFResponse, error) {
	return &cpf.GenerateCPFResponse{
		Result: s.Generator.Generate(),
	}, nil
}
