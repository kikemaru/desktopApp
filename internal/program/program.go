package program_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-playground/validator/v10"
	funeapplication "github.com/kikemaru/desktopApp/pkg/Fune"
	"github.com/kikemaru/desktopApp/pkg/logger"
	validatorpackage "github.com/kikemaru/desktopApp/pkg/validator"
)

type Application struct {
	fune  fyne.App
	valid *validator.Validate
	log   *logger.ZerologLogger
	uc    AppUC
}

func NewApp(fune *funeapplication.App, v *validatorpackage.Validator, logger *logger.ZerologLogger, uc AppUC) *Application {
	return &Application{
		fune:  *fune.A,
		valid: v.V,
		log:   logger,
		uc:    uc,
	}
}

func (app *Application) MainWindow() {
	w := app.fune.NewWindow("Hello")

	text := app.uc.GetDataForMyApp()
	hello := widget.NewLabel(text)
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
