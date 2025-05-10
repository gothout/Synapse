package validators

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateCNPJ(cnpj string) error {
	cnpj = strings.TrimSpace(cnpj)
	re := regexp.MustCompile(`[^0-9]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return fmt.Errorf("CNPJ deve conter exatamente 14 dígitos")
	}

	todosIguais := true
	for i := 1; i < 14 && todosIguais; i++ {
		if cnpj[i] != cnpj[0] {
			todosIguais = false
		}
	}
	if todosIguais {
		return fmt.Errorf("CNPJ inválido: todos os dígitos são iguais")
	}

	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	soma := 0
	for i := 0; i < 12; i++ {
		soma += int(cnpj[i]-'0') * pesos1[i]
	}
	resto := soma % 11
	dv1 := 0
	if resto >= 2 {
		dv1 = 11 - resto
	}

	soma = 0
	for i := 0; i < 13; i++ {
		soma += int(cnpj[i]-'0') * pesos2[i]
	}
	resto = soma % 11
	dv2 := 0
	if resto >= 2 {
		dv2 = 11 - resto
	}

	if int(cnpj[12]-'0') != dv1 || int(cnpj[13]-'0') != dv2 {
		return fmt.Errorf("CNPJ inválido: dígitos verificadores incorretos")
	}

	return nil
}
