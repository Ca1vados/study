package usecase

import "temp/config"

type UseCase struct {
	LogLvl string
}

func New(cfg *config.Config) *UseCase {
	u := UseCase{
		LogLvl: cfg.LogLvl,
	}
	return &u
}
