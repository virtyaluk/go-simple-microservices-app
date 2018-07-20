package routers

import (
  "github.com/gorilla/mux"
  "github.com/virtyaluk/go-simple-microservices-app/bookings/controllers"
)

func SetBookingsRouters(router *mux.Router) *mux.Router {
  router.HandleFunc("/bookings", controllers.GetBookings).Methods("GET")
  router.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")

  return router
}
