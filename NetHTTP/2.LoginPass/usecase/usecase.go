package usecase

import (
	"errors"
	"fmt"
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
	u.db.CreateUser(user)
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
		u.db.CreateUser(user)
		return nil
	} else {
		return errors.New("logins is not free ...")
	}

}

func (u *UseCase) Login(user entity.User) (entity.User, error) {
	//в функцию поступает структура user с заполненными полями login и passhash
	// нужно запустить функцию getuser с предложенным логином.
	// при успехе сверить пароли, при успехе записать в user секрет.
	userFromDb, err := u.db.GetUser(user.Login)
	if err != nil {
		fmt.Errorf("user %s is not available", user.Login)
	}
	if user.PassHash {

	}

	return user, err
}

func (u *UseCase) GetAllUsers() ([]entity.User, error) {
	// 1. Валидация входных данных - если нужно

	// Логика приложения
	users, err := u.db.GetAllUsers()
	if err != nil {
		e := fmt.Errorf("Что-то пошло не так в \"users, err := u.db.GetAllUsers()\": %v", err)
		return nil, e
	}

	return users, nil
}
