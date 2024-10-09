package main

import (
	"github.com/maximianodev/scheduler/routes"
)

func main() {
	router := routes.Routes()

	router.Run("localhost:8080")
}
