package app

import (
	program_window "github.com/kikemaru/desktopApp/internal/program"
	getDataRepo "github.com/kikemaru/desktopApp/internal/repositories/mssql/getData"
	GetData_uc "github.com/kikemaru/desktopApp/internal/usecases/GetData"
	"github.com/kikemaru/desktopApp/pkg/Fune"
	"github.com/kikemaru/desktopApp/pkg/logger"
	"github.com/kikemaru/desktopApp/pkg/validator"
	"log"
)

func Run() {
	application := Fune.InitApplication()
	validator := validator.InitValidator()
	l, err := logger.NewZerologLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}

	repo := getDataRepo.NewRepo(l)
	usecase := GetData_uc.NewUC(repo, l)
	program := program_window.NewApp(application, validator, l, usecase)
	program.MainWindow()

}
