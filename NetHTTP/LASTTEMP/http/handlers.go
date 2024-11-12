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
	data, err := hs.u.GetGetCryptoRatesBinance()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}
