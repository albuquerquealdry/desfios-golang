package main

import (
	"fmt"
)

type Payment interface {
	Pay(value int64)
}

type CourrentAccount struct {
	id      string
	balance int64
}

type CriptoWallet struct {
	hash    string
	balance int64
}

func (cw *CriptoWallet) Pay(value int64) {
	cw.balance -= value
	fmt.Printf("The current value in wallet %s is R$ %s ", cw.hash, cw.balance)
}

func (ca CourrentAccount) Pay(value int64) {
	ca.balance -= value
	fmt.Printf("The current value in Courrent Account %s is R$ %.2f ", ca.id, ca.balance)
}

func main() {
	crypto := &CriptoWallet{
		hash:    "218903890283190832dsadasds",
		balance: 1000,
	}

	crypto.Pay(200)
	fmt.Println(crypto.balance)
}
