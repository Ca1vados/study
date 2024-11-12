package usecase

import (
	"fmt"
	"temp/config"
	"temp/database"
)

type UseCase struct {
	LogLvl string
	db     *database.DataBase
}

func New(db *database.DataBase, cfg *config.Config) *UseCase {
	return &UseCase{
		LogLvl: cfg.LogLvl,
		db:     db,
	}
}

func (u *UseCase) GetData() string {
	fmt.Printf("[%s] (usecase) Вызвали функцию GetData\n", u.LogLvl)
	data := u.db.GetData()
	return data
}
