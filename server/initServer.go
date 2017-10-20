package server

import (
	"net/http"
)

func InitServer(Port string, H http.Handler) {

	http.ListenAndServe(Port,H)
}