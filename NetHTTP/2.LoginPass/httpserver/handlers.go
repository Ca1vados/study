package httpserver

import (
	"encoding/json"
	"loginpass/entity"
	"net/http"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// Register принимает в body логин, пароль, секрет  - в ответе либо OK либо error
// @Summary добавляет пользователя
// @Description принимает в body логин, пароль, секрет  - в ответе либо OK либо error
// @Tags register
// @Accept json
// @Produce json
// @Param login_pass_secret body entity.UserReg true "форма для добавления пользователя"
// @Success 200 {int} http.StatusCreated
// @Router /register [post]
func (hs *HttpServer) Register(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	var user entity.UserReg
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err := hs.u.RegisterUser(user)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
}

// Login на запрос по логину и паролю возвращает секрет
// @Summary authorisation by login password
// @Description а запрос по логину и паролю возвращает секрет из базы или оишбку если логин или пароль неверны
// @Tags authorisation
// @Accept json
// @Produce json
// @Param authorisationrequest body entity.User true "форма авторизации"
// @Success 200 {array} entity.User
// @Router /login [post]
func (hs *HttpServer) Login(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	userWithSecret, err := hs.u.Login(user)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userWithSecret)
}

// GetAllUsers возвращает все данные из базы
// @Summary возвращает все данные из базы
// @Description при запросе возвращает все данные из базы данных
// @Tags database
// @Produce json
// @Success 200 {array} entity.User
// @Router /getallusers [get]
func (hs *HttpServer) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// 1. Обработка входных данных и продготовка данных для передачи в UseCase
	// тут нет обработки request, так как это просто запрос данных

	// 2. Вызов метода UseCase - внутри вся логика
	enableCORS(w)
	users, err := hs.u.GetAllUsers() // вызов UseCase

	// 3. Отправка ответа клиенту
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func (hs *HttpServer) Hello(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Write([]byte("Hello!"))
}
