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

	intCode, _ := strconv.ParseInt(code, 10, 0)
	w.WriteHeader(int(intCode))
	w.Write(message)

	return w
}
