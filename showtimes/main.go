package main

import (
  "net/http"
  "log"
  "github.com/virtyaluk/go-simple-microservices-app/showtimes/common"
  "github.com/virtyaluk/go-simple-microservices-app/showtimes/routers"
)

// Entry point for the program
func main() {
  // Calls startup logic
  common.StartUp()
  // Get the mus router object
  router := routers.InitRoutes()

  server := &http.Server{
    Addr: common.AppConfig.Server,
    Handler: router,
  }

  log.Println("Listening...")
  server.ListenAndServe()
}
