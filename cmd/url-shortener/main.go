package main

import (
	"fmt"
	"github.com/sdvaanyaa/url-shortener/internal/config"
)

func main() {
	// init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, chi render

	// TODO: run server
}
