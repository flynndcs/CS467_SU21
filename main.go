package main

import (
	"os"

	"flynndcs.com/flynndcs/grpc-gateway/src/environment"
)

func main() {
	environment.RegisterAndServeEnvironment(os.Args[1], os.Args[2])
}
