// Package classification Account API.
//
// this is to show how to write RESTful APIs in golang.
// that is to provide a detailed overview of the language specs
//
// Terms Of Service:
//
//     Schemes: https
//     Host: nsr-ormg-test.sodimac-it.com/nsrorffmsdecryptqrsod
//     Version: 1.0.0
//     Contact: Gabriela Loreto <mariagabriela191@hotmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: API-KEY
//          in: header
//
// swagger:meta
package main

import (
	"app/src/env"
	router "app/src/http"
	"app/src/server"
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	fs := http.FileServer(http.Dir("./swaggerui"))

	// Declare bundle server
	server.HealthCheckRoute(httpRouter)
	server.GetTopSecretRoute(httpRouter)

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.SWAGGERUI(fs)

	// UP server
	httpRouter.SERVE(env.PORT_SERVER)
}
