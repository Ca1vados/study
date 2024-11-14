package http

import (
	"encoding/json"
	"net/http"
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
	data, err := hs.u.GetAndSaveCryptoRatesBinance()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}
