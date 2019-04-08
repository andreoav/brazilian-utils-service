package server

import (
	"net"

	generator "github.com/andreoav/brazilian-utils-service/pkg/generators/brutils"
	validator "github.com/andreoav/brazilian-utils-service/pkg/validators/brutils"

	"github.com/andreoav/brazilian-utils-service/internal/transport"
	"github.com/andreoav/brazilian-utils-service/pkg/api/cpf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() error {
	listen, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		return err
	}

	s := grpc.NewServer()

	cpf.RegisterGenerateCPFServiceServer(s, &transport.GenerateCPFService{
		Generator: &generator.BrazilianUtilsCPFGenerator{},
	})

	cpf.RegisterValidateCPFServiceServer(s, &transport.ValidateCPFService{
		Validator: &validator.BrazilianUtilsCPFValidator{},
	})

	reflection.Register(s)

	return s.Serve(listen)
}
