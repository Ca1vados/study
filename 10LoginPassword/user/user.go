package user

type User struct {
	Login    string `json:"Id"`
	PassHash string `json:""`
	Secret   string `json:"Secret"`
}

/*
будет принимать введенные от пользователя данные,
сверять их с данными полученными функцией Database
и при удачной верификации возвращать дополненную полем секрет структуру
*/
