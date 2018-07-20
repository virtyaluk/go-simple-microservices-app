package controllers

import (
  "net/http"
  "encoding/json"
  "gopkg.in/mgo.v2"
  "github.com/gorilla/mux"
  "github.com/virtyaluk/go-simple-microservices-app/showtimes/common"
  "github.com/virtyaluk/go-simple-microservices-app/showtimes/data"
)

// Handler for HTTP Get - "/showtimes"
// Returns all Showtime documents
func GetShowTimes(responseWriter http.ResponseWriter, request *http.Request) {
  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("showtimes")
  repo := &data.ShowTimeRepository{c}
  // Get all showtimes form repository
  showtimes := repo.GetAll()
  j, err := json.Marshal(ShowTimesResource{Data: showtimes})

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  // Send response back
  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}

// Handler for HTTP Post - "/showtimes"
// Create a new Showtime document
func CreateShowTime(responseWriter http.ResponseWriter, request *http.Request) {
  var dataResource ShowTimeResource
  // Decode the incoming ShowTime json
  err := json.NewDecoder(request.Body).Decode(&dataResource)

  if err != nil {
    common.DisplayAppError(responseWriter, err, "Invalid ShowTime data", 500)

    return
  }

  showtime := &dataResource.Data
  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("showtimes")
  // Create ShowTime
  repo := &data.ShowTimeRepository{c}
  repo.Create(showtime)
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

// Handler for HTTP Get - "/shotimes/{date}"
// Return ShowTime document
func GetShowTimeByDate(responseWriter http.ResponseWriter, request *http.Request) {
  // Get date from incoming url
  vars := mux.Vars(request)
  date := vars["date"]

  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("showtimes")
  repo := &data.ShowTimeRepository{c}

  // Get showtime by date
  showtime, err := repo.GetByDate(date)

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  // Create data for the response
  j, err := json.Marshal(ShowTimeResource{Data: showtime})

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  // Send response back
  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}

// Handler for HTTP Delete - "/showtimes/{id}"
// Delete a Showtime document by id
func DeleteShowTime(responseWriter http.ResponseWriter, request *http.Request) {
  // Get id from incoming url
  vars := mux.Vars(request)
  id := vars["id"]

  // Create new context
  context := NewContext()
  defer context.Close()


  c := context.DbCollection("showtimes")
  // Remove showtime by id
  repo := &data.ShowTimeRepository{c}
  err := repo.Delete(id)

  if err != nil {
    if err == mgo.ErrNotFound {
      responseWriter.WriteHeader(http.StatusNotFound)

      return
    } else {
      common.DisplayAppError(responseWriter, err, "An unexpected error ahs occurred", 500)

      return
    }
  }

  // Send response back
  responseWriter.WriteHeader(http.StatusNoContent)
}
