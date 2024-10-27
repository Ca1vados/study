package usecase

import (
	"errors"
	"loginpass/database"
	"loginpass/entity"
)

type UseCase struct {
	// какие-то поля
	db *database.DataBase
}

func New(data *database.DataBase) *UseCase {
	u := UseCase{db: data}
	return &u
}

func (u *UseCase) RegisterUser(user entity.User) error {

	allLogins, err := u.db.GetAllLogins()
	if err != nil {
		return err
	}

	isFree := true
	for _, login := range allLogins {
		if login == user.Login {
			isFree = false
			break
		}
	}

	if isFree {
		u.db.AddUser(user)
		return nil
	} else {
		return errors.New("logins is not free ...")
	}

}

func (u *UseCase) Login(user entity.User) error {
	return nil
}
