package GetData_uc

import "github.com/kikemaru/desktopApp/pkg/logger"

type UC struct {
	repo Repo
	log  *logger.ZerologLogger
}

func NewUC(r Repo, l *logger.ZerologLogger) *UC {
	return &UC{
		repo: r,
		log:  l,
	}
}

func (uc *UC) GetDataForMyApp() string {
	return uc.repo.GetDataForMyApp()
}
