package main

func main() {

	oAuthProcessor := &OAuthProcessor{}

	usernamePasswordProcessor := &UsernamePasswordProcessor{}
	usernamePasswordProcessor.SetNext(oAuthProcessor)

	provider := AuthenticationProvider{"O_AUTH_PROVIDER"}

	usernamePasswordProcessor.IsAuthorized(provider)

}
