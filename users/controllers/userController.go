package controllers

import (
  "net/http"
  "encoding/json"
  "gopkg.in/mgo.v2"
  "github.com/gorilla/mux"
  "github.com/virtyaluk/go-simple-microservices-app/users/data"
  "github.com/virtyaluk/go-simple-microservices-app/users/common"
)

// Handler for HTTP Get - "/users"
// Returns all User documents
func GetUsers(responseWriter http.ResponseWriter, request *http.Request) {
  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("users")
  repo := &data.UserRepository{c}
  // Get all users form repository
  users := repo.GetAll()
  j, err := json.Marshal(UsersResource{Data: users})

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  // Send response back
  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}

// Handler for HTTP Post - "/users"
// Create a new Showtime document
func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
  var dataResource UserResource
  // Decode the incoming User json
  err := json.NewDecoder(request.Body).Decode(&dataResource)

  if err != nil {
    common.DisplayAppError(responseWriter, err, "Invalid User data", 500)

    return
  }

  user := &dataResource.Data
  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("users")
  // Create User
  repo := &data.UserRepository{c}
  repo.Create(user)
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

// Handler for HTTP Delete - "/users/{id}"
// Delete a User document by id
func DeleteUser(responseWriter http.ResponseWriter, request *http.Request) {
  // Get id from incoming url
  vars := mux.Vars(request)
  id := vars["id"]

  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("users")
  // Remove user by id
  repo := &data.UserRepository{c}
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
