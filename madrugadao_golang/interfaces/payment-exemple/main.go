package main

import "fmt"

// Definindo a interface ContaBancaria
type ContaBancaria interface {
	Depositar(valor float64) // Método para depositar
	Sacar(valor float64)     // Método para sacar
	Saldo() float64          // Método para verificar saldo
}

// Estrutura ContaCorrente
type ContaCorrente struct {
	Titular string
	saldo   float64
}

// Estrutura ContaPoupanca
type ContaPoupanca struct {
	Titular string
	saldo   float64
}

// Implementando o método Depositar para ContaCorrente
func (c *ContaCorrente) Depositar(valor float64) {
	c.saldo += valor
	fmt.Println("Depósito de R$", valor, "realizado na Conta Corrente")
}

// Implementando o método Sacar para ContaCorrente
func (c *ContaCorrente) Sacar(valor float64) {
	if valor > c.saldo {
		fmt.Println("Saldo insuficiente na Conta Corrente")
	} else {
		c.saldo -= valor
		fmt.Println("Saque de R$", valor, "realizado da Conta Corrente")
	}
}

// Implementando o método Saldo para ContaCorrente
func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}

// Implementando o método Depositar para ContaPoupanca
func (c *ContaPoupanca) Depositar(valor float64) {
	c.saldo += valor
	fmt.Println("Depósito de R$", valor, "realizado na Conta Poupança")
}

// Implementando o método Sacar para ContaPoupanca
func (c *ContaPoupanca) Sacar(valor float64) {
	if valor > c.saldo {
		fmt.Println("Saldo insuficiente na Conta Poupança")
	} else {
		c.saldo -= valor
		fmt.Println("Saque de R$", valor, "realizado da Conta Poupança")
	}
}

// Implementando o método Saldo para ContaPoupanca
func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}

// Função que processa operações em qualquer tipo de conta
func ProcessarOperacao(c ContaBancaria, operacao string, valor float64) {
	switch operacao {
	case "depositar":
		c.Depositar(valor)
	case "sacar":
		c.Sacar(valor)
	case "saldo":
		fmt.Println("Saldo da conta:", c.Saldo())
	default:
		fmt.Println("Operação desconhecida")
	}
}

func main() {
	// Criando instâncias de ContaCorrente e ContaPoupanca
	contaCorrente := &ContaCorrente{Titular: "João", saldo: 1000.00}
	contaPoupanca := &ContaPoupanca{Titular: "Maria", saldo: 1500.00}

	// Realizando operações
	ProcessarOperacao(contaCorrente, "depositar", 200.00) // Depósito na Conta Corrente
	ProcessarOperacao(contaPoupanca, "sacar", 500.00)     // Saque na Conta Poupança
	ProcessarOperacao(contaCorrente, "saldo", 0)          // Verificar saldo da Conta Corrente
	ProcessarOperacao(contaPoupanca, "saldo", 0)          // Verificar saldo da Conta Poupança
	ProcessarOperacao(contaPoupanca, "sacar", 2000.00)    // Tentando sacar mais que o saldo
}
