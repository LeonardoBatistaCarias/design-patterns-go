package main

type AuthenticationProcessor interface {
	IsAuthorized(authProvider AuthenticationProvider) bool
	SetNext(nextProcessor AuthenticationProcessor)
}
