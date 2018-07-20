package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
  router := mux.NewRouter().StrictSlash(false)
  // Routes for the Booking entity
  router = SetBookingsRouters(router)

  return router
}
