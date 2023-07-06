package main

import (
	"log"

	"github.com/neoetheilred/crispy-pancake/common"
)

func main() {
	log.Println("Starting server...")
	common.StartBookApi()
}
