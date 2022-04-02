package main

func main() {
	bankAccount := NewBankAccount()
	bankAccount.
		WithName("Leonardo").
		WithAccountNumber("123456789-0").
		WithEmail("leonardo@email.com").
		wantNewsLetter(true)

	bankAccount.toString()
}
