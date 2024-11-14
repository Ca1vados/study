package usecase

import (
	"temp/config"
	"temp/database"
)

type UseCase struct {
	logLvl string
	db     *database.DataBase
}

func New(db *database.DataBase, cfg *config.Config) *UseCase {
	usecase := UseCase{
		logLvl: cfg.LogLvl,
		db:     db,
	}
	return &usecase
}
