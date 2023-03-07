package repo

import (
	"database/sql"
	"fmt"
)

type IAnimalRepo interface {
	DoSomething(test string) (string, error)
}

type AnimalRepo struct {
	DB *sql.DB
}

func (a *AnimalRepo) DoSomething(test string) (string, error) {
	val := a.DB.Stats()
	sVal := fmt.Sprintf("%f", val.WaitDuration.Hours())

	return sVal, nil
}
