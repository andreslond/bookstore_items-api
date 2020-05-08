package controller

import "net/http"

const (
	pong = "pong"
)

var (
	PingController PingControllerInterface = &pingControllers{}
)

type PingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingControllers struct{}

func (u *pingControllers) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
