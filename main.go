package main

import "payments-go/infrastructure/http/server"

func main() {
	server.
		NewConfig().
		WithAppConfig().
		InitLogger().
		WithDB().
		WithCache().
		WithRepository().
		WithWebServer().
		Start()
}
