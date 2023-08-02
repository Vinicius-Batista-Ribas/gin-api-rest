package main

import (
	"conta"
	"fmt"
)

func PagarBoleto(conta verificaConta, valor float64) {
	conta.sacar(valor)
}

type verificaConta interface {
	sacar(valor float64) string
}

func main() {
	conta1 := conta.ContaCorrente{}
	conta1.Depositar(100)
	fmt.Println(conta1)

}
