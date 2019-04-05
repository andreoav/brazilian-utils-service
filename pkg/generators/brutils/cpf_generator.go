package brutils

import (
	"github.com/brazilian-utils/go-brazilian-utils/cpf"
)

type BrazilianUtilsCPFGenerator struct{}

func (*BrazilianUtilsCPFGenerator) Generate() string {
	return cpf.Generate()
}
