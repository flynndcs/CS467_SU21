package main

import (
	"os"

	"CS467_SU21/src/environment"
)

func main() {
	environment.RegisterAndServeEnvironment(os.Args[1], os.Args[2])
}
