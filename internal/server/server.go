package server

import (
	"fmt"
	"net"
	"os"

	"github.com/andreoav/brazilian-utils-service/pkg/generators/brutils"

	"github.com/andreoav/brazilian-utils-service/internal/transport"
	"github.com/andreoav/brazilian-utils-service/pkg/api/cpf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("PORT")))

	if err != nil {
		return err
	}

	s := grpc.NewServer()

	// cpf.RegisterFormatServiceServer(s, &transport.FormatterService{
	// 	Formatter: &formatter.CPFFormatter{},
	// })

	cpf.RegisterGenerateCPFServiceServer(s, &transport.GenerateCPFService{
		Generator: &brutils.BrazilianUtilsCPFGenerator{},
	})

	reflection.Register(s)

	return s.Serve(listen)
}
