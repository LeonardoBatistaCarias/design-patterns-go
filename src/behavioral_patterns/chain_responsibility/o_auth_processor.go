package main

import "fmt"

type OAuthProcessor struct {
	next AuthenticationProcessor
}

func (o *OAuthProcessor) IsAuthorized(provider AuthenticationProvider) bool {
	fmt.Println("Checking if is OAuthProvider")
	if provider.GetTypeAuthProvider() == "O_AUTH_PROVIDER" {
		fmt.Println("It's OAuth Provider")
		return true
	} else if o.next != nil {
		o.next.IsAuthorized(provider)
	}
	return false
}

func (o *OAuthProcessor) SetNext(next AuthenticationProcessor) {
	o.next = next
}
