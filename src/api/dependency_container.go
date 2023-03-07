package api

import (
	"database/sql"

	ctl "github.com/TimetoPretend54/go-chatgpt-copilot/api/controller"
	"github.com/TimetoPretend54/go-chatgpt-copilot/repo"
	"github.com/TimetoPretend54/go-chatgpt-copilot/service"
)

type DependencyContainer struct {
	DogCtl *ctl.DogController
	CatCtl *ctl.CatController
}

func InitDependencies() DependencyContainer {
	animalRepo := &repo.AnimalRepo{DB: &sql.DB{}}
	animalService := &service.AnimalService{Repo: animalRepo}

	return DependencyContainer{
		DogCtl: &ctl.DogController{Service: animalService},
		CatCtl: &ctl.CatController{Service: animalService},
	}
}
