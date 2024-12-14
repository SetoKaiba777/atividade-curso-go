package main

import (
	"api-pedidos/infrastructure/setup"
	"context"
	"sync"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	setup.
		NewConfig().
		InitLogger().
		WithAppConfig().
		WithDB().
		WithRouter().
		WithWebServer().
		Start(ctx, &wg)

	wg.Wait()
}
