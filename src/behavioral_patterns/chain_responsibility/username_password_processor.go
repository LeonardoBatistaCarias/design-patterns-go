package main

import "fmt"

type UsernamePasswordProcessor struct {
	next AuthenticationProcessor
}

func (u *UsernamePasswordProcessor) IsAuthorized(provider AuthenticationProvider) bool {
	fmt.Println("Checking if is UsernamePasswordProvider")
	if provider.GetTypeAuthProvider() == "USERNAME_PASSWORD_PROVIDER" {
		fmt.Println("It's UsernamePassword Provider")
		return true
	} else if u.next != nil {
		u.next.IsAuthorized(provider)
	}
	return false
}

func (u *UsernamePasswordProcessor) SetNext(next AuthenticationProcessor) {
	u.next = next
}
