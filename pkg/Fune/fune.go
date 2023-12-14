package Fune

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	A *fyne.App
}

func InitApplication() *App {
	application := app.New()

	return &App{A: &application}
}
