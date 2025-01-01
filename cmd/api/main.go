package main

import (
	"fmt"
	"net/http"

	"github.com/Ibrahimkazi18/First-Go/internal/handler"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handler.Handler(r)

	fmt.Println("Starting Go API service....")

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
