package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/virtyaluk/go-simple-microservices-app/bookings/data"
  "github.com/virtyaluk/go-simple-microservices-app/bookings/common"
)

// Handler for HTTP Post - "/bookings"
// Create a new Booking document
func CreateBooking(responseWriter http.ResponseWriter, request *http.Request) {
  var dataResource BookingResource
  // Decode the incoming Booking json
  err := json.NewDecoder(request.Body).Decode(&dataResource)

  if err != nil {
    common.DisplayAppError(responseWriter, err, "Invalid Booking data", 500)

    return
  }
  booking := &dataResource.Data
  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("bookings")
  // Create Booking
  repo := &data.BookingRepository{c}
  repo.Create(booking)
  // Create response data
  j, err := json.Marshal(dataResource)

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  // Send response back
  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}

// Handler for HTTP Get - "/bookings"
// Return all Booking documents
func GetBookings(responseWriter http.ResponseWriter, request *http.Request) {
  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("bookings")
  repo := &data.BookingRepository{c}
  // Get all bookings
  bookings := repo.GetAll()
  // Create response data
  j, err := json.Marshal(BookingsResource{Data: bookings})

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  // Send response back
  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}
