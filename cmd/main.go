package main

import (
	"github.com/kikemaru/desktopApp/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("cannot start application: %v", err)
	}
}
