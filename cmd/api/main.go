package main

import (
	"fmt"
	"net/http"

	// To install extern modules -> go mod tidy
	"github.com/eggermarc/go-api/internal/handlers"
	"github.com/go-chi/chi"          // web dev package
	log "github.com/sirupsen/logrus" // log errors
)

func main() {
   log.SetReportCaller(true)
   var r *chi.Mux = chi.NewRouter()
   handlers.Handler(r)

   fmt.Println("Welcome to Go API")
   fmt.Println(`
   ______     ______        ______     ______   __    
  /\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
  \ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
   \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
    \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)
   err := http.ListenAndServe("localhost:8000", r)
   if err != nil {
       log.Error(err)
   }

}
