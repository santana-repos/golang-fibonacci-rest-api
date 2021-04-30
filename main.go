package main

import "go-fib/api"

func main() {

	server := api.NewApiServer(8000)
	server.StartApiServer()

}
