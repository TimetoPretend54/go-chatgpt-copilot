package service

import (
	"fmt"

	"github.com/TimetoPretend54/go-chatgpt-copilot/repo"
)

type IAnimalService interface {
	DoSomething(test string) (string, error)
}

type AnimalService struct {
	Repo repo.IAnimalRepo
}

func (a *AnimalService) DoSomething(test string) (string, error) {
	val, err := a.Repo.DoSomething("test")
	if err != nil {
		return "", fmt.Errorf("error running animalRepo: %w", err)
	}

	return val, nil
}
