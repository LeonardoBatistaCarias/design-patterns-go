package main

import "fmt"

type BankAccountBuilder struct {
	Name          string
	AccountNumber string
	Email         string
	NewsLetter    bool
}

func NewBankAccount() *BankAccountBuilder {
	return &BankAccountBuilder{}
}

func (b *BankAccountBuilder) WithName(name string) *BankAccountBuilder {
	b.Name = name
	return b
}

func (b *BankAccountBuilder) WithAccountNumber(accountNumber string) *BankAccountBuilder {
	b.AccountNumber = accountNumber
	return b
}

func (b *BankAccountBuilder) WithEmail(email string) *BankAccountBuilder {
	b.Email = email
	return b
}

func (b *BankAccountBuilder) wantNewsLetter(newsLetter bool) *BankAccountBuilder {
	b.NewsLetter = newsLetter
	return b
}

func (b *BankAccountBuilder) toString() {
	fmt.Println("Bank Account")
	fmt.Printf("Name: %s, Account Number: %s, Email: %s, News Letter: %t.", b.Name, b.AccountNumber, b.Email, b.NewsLetter)
}
