package http

import (
	"encoding/json"
	"fmt"
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
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (hs *HttpServer) History(w http.ResponseWriter, r *http.Request) {
	data, err := hs.u.GetHistory()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (hs *HttpServer) Convert(w http.ResponseWriter, r *http.Request) {
	var newConvert entity.ConvertRequest
	err := json.NewDecoder(r.Body).Decode(&newConvert)
	if err != nil {
		err = fmt.Errorf("bad request: %v", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	data, err := hs.u.Convert(newConvert)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}
