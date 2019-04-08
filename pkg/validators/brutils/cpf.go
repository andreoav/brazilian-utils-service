package brutils

import (
	"github.com/brazilian-utils/go-brazilian-utils/cpf"
)

type BrazilianUtilsCPFValidator struct{}

func (*BrazilianUtilsCPFValidator) IsValid(input string) bool {
	return cpf.IsValid(input)
}
