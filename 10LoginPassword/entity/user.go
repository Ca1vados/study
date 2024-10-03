package entity

type User struct {
	Login    string `json:"login"`
	Pass     string `json:"pass"` // убрать из структуры и из базы
	PassHash string `json:"hash"`
	Secret   string `json:"secret"`
}

/*
будет принимать введенные от пользователя данные,
сверять их с данными полученными функцией Database
и при удачной верификации возвращать дополненную полем секрет структуру
*/
func (u *User) VeryficationPass(pass string) bool {
	// измени на проверку - u.PassHash == sha256(pass)
	return pass == u.Pass
}
