package http

import (
	"encoding/json"
	"net/http"
	"temp/entity"
)

func (hs *HttpServer) Hello(w http.ResponseWriter, r *http.Request) {
	data := "Hello!"
	json.NewEncoder(w).Encode(data)
}

// GetCryptoRates получает курсы криптовалют и сохраняет их в бд (технический комментарий)
// @Summary Get rates
// @Description получает курсы криптовалют и сохраняет их в бд (комментарий для пользователя)
// @Tags cryptorates
// @Accept json
// @Produce json
// @Success 200 {array} entity.ApiBinanceResponse
// @Router /get [get]
func (hs *HttpServer) GetCryptoRataes(w http.ResponseWriter, r *http.Request) {
	data, err := hs.u.GetAndSaveCryptoRatesBinance()
	//data, err := hs.u.GetCryptoRatesBinance()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(data)
}

// History возвращает историю хранимых курсов валют
// @Summary Get history
// @Description возвращает историю хранимых курсов валют
// @Tags cryptorates
// @Produce json
// @Success 200 {array} entity.ApiBinanceResponse
// @Router /history [get]
func (hs *HttpServer) History(w http.ResponseWriter, r *http.Request) {
	data, err := hs.u.GetHistory()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Convert возвращает ответ на запрос по вычислению криптовалют
// @Summary Get response
// @Description возвращает ответ на запрос по вычислению криптовалют
// @Tags convert
// @Accept json
// @Produce json
// @Param convertrequest body entity.ConvertRequest true "форма для ковертации"
// @Success 200 {array} entity.ConvertResponse
// @Router /convert [post]

// параметры: название параметра, куда принимает, форма приемки, обязателен ли параметр, название формы параметра
func (hs *HttpServer) Convert(w http.ResponseWriter, r *http.Request) {
	var newConvert entity.ConvertRequest
	err := json.NewDecoder(r.Body).Decode(&newConvert)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	//data, err := hs.u.Convert(newConvert)
	data, err := hs.u.Convert2(newConvert)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(data)
}
