package main

import (
	"github.com/Nachofra/final-esp-backend-3/cmd/api/config"
)

func main() {
	_, err := config.Get()
	if err != nil {
		panic(err)
	}
}
