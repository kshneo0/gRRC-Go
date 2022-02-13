package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello world...")
	log.Info().Msg("Hi")
}
