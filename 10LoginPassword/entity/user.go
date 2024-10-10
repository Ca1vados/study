package entity

import "crypto/sha256"

type User struct {
	Login    string   `json:"login"`
	PassHash [32]byte `json:"hash"` // поменял пароль в структуре на хэш, поменял тип (нужно пояснение по типу byte и [32]byte)
	Secret   string   `json:"secret"`
}

/*
будет принимать введенные от пользователя данные,
сверять их с данными полученными функцией Database
и при удачной верификации возвращать дополненную полем секрет структуру
*/

func (u *User) VeryficationPass(pass string) bool {
	// измени на проверку - u.PassHash == sha256(pass)
	inputPassHash := sha256.Sum256([]byte("pass")) // pass поступающий в функцию перевел в хэш и на вывод поставил сравнение с PassHash
	return inputPassHash == u.PassHash
}
