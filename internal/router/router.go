package router

import (
	"golang-training/main/internal/sms"

	"github.com/go-chi/chi/v5"
	// "github.com/go-delve/delve/service"
)

const (
	pathGetData = "/sms/Getdata"
	pathAddData = "/sms/Adddata"
)

func InitRouter(service sms.Service)  *chi.Mux{
	r := chi.NewRouter()
	r.Get(pathGetData, service.GetData)
	r.Post(pathAddData, service.AddData)
	return r
}