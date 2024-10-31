package entity

type User struct {
	Login    string `json:"login"`
	PassHash string `json:"hash"` // поменял пароль в структуре на хэш, поменял тип (нужно пояснение по типу byte и [32]byte)
	Secret   string `json:"secret"`
}
