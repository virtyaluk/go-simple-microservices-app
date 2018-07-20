package routers

import (
  "github.com/gorilla/mux"
  "github.com/virtyaluk/go-simple-microservices-app/users/controllers"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
  router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
  router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
  router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

  return router
}