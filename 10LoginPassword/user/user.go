package user

type User struct {
	Login    string `yaml:"Id"`
	Pass     string `json:"Pass"`
	PassHash string `json:""`
	Secret   string `json:"Secret"`
}
