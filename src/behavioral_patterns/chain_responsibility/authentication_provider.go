package main

type AuthenticationProvider struct {
	typeAuthProvider string
}

func (a *AuthenticationProvider) GetTypeAuthProvider() string {
	return a.typeAuthProvider
}
