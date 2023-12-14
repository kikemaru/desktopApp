package getDataRepo

import "github.com/kikemaru/desktopApp/pkg/logger"

type Repo struct {
	logger *logger.ZerologLogger
}

func NewRepo(logger *logger.ZerologLogger) *Repo {
	return &Repo{
		logger: logger,
	}
}

func (rp *Repo) GetDataForMyApp() string {
	return "123testapp hello world!"
}
