package usecase

import (
	"fmt"
	"temphttp/config"
	"temphttp/database"
)

type Usecase struct {
	log_level string
	db        *database.Database
}

func NewUsecase(db *database.Database, cfg *config.Config) *Usecase {
	return &Usecase{
		log_level: cfg.LogLevel,
		db:        db,
	}
}

func (u *Usecase) GetData() string {
	fmt.Printf("[%s] (usecase) Вызвали функцию GetData\n", u.log_level)
	data := u.db.GetData()
	return data
}

func (u *Usecase) PutData(data string) {
	fmt.Printf("[%s]  (usecase) Вызвали функцию PutData\n", u.log_level)
	u.db.CreateData(data)
}
