package handlers

import (
	"Ys7/config"
	"encoding/json"
	"net/http"
	"strconv"
)

func ErrorHandler(w http.ResponseWriter, msg string, code string) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")

	errJson := config.Error{
		Msg:  msg,
		Code: code,
	}
	message, _ := json.Marshal(errJson)

	intCode, _ := strconv.Atoi(code)
	w.WriteHeader(intCode)
	w.Write(message)

	return w
}
