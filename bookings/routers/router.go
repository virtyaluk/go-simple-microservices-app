package routers

import "github.com/gorilla/mux"

func InitRouts() *mux.Router {
  router := mux.NewRouter().StrictSlash(false)
  // Routes for the Booking entity
  router = SetBookingsRouters(router)

  return router
}
