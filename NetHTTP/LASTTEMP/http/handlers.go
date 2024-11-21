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

func (hs *HttpServer) GetCryptoRataes(w http.ResponseWriter, r *http.Request) {
	data, err := hs.u.GetAndSaveCryptoRatesBinance()
	//data, err := hs.u.GetCryptoRatesBinance()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (hs *HttpServer) History(w http.ResponseWriter, r *http.Request) {
	data, err := hs.u.GetHistory()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(data)
}

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
