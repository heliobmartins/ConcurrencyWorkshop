package main

import (
	"github.com/heliobmartins/concurrencyworkshop/internal/challenge/implme/intermediate/poke/app"
	"github.com/heliobmartins/concurrencyworkshop/internal/challenge/implme/intermediate/poke/client"
)

func main() {
	pokeAPP := app.NewPokeApp(client.New())
	pokeAPP.Start()
}
