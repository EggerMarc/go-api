package main

import (
    "fmt"
    "net/http"

    // To install extern modules -> go mod tidy
    "github.com/go-chi/chi" // web dev package
    "github.com/eggermarc/go-api/internal/handlers" // internal package?
    log "github.com/sirupsen/logrus" // log errors
)

func main() {
   log.SetReportCaller(true)
   var r *chi.Mux = chi.NewRouter()
   handlers.Handler(r)

   fmt.Println("Welcome to Go API")
   err := http.ListenAndServe("localhost:800", r)
   if err != nil {
       log.Error(err)
   }

}
