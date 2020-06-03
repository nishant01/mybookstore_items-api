package controllers

import (
	"net/http"
)

const (
	pong = "pong"
)

var (
	PingController PingControllerInterface = &pingController{}
)

type PingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct {

}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
