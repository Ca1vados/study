package user

type User struct {
	Login    string `yaml:"Id"`
	PassHash string `json:""`
	Secret   string `json:"Secret"`
}
