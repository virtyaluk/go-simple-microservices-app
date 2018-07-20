package controllers

import (
  "net/http"
  "encoding/json"
  "gopkg.in/mgo.v2"
  "github.com/gorilla/mux"
  "github.com/virtyaluk/go-simple-microservices-app/movies/common"
  "github.com/virtyaluk/go-simple-microservices-app/movies/data"
)

// Handler for HTTP Get - "/movies"
// Return all Movie documents
func GetMovies(responseWriter http.ResponseWriter, request *http.Request) {
  context := NewContext()
  defer context.Close()

  c:= context.DbCollection("movies")
  repo := &data.MovieRepository{c}
  movies := repo.GetAll()
  j, err := json.Marshal(MoviesResource{Data: movies})

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}

// Handler for HTTP post - "/movies"
// Insert a new movie document
func CreateMovie(responseWriter http.ResponseWriter, request *http.Request) {
  var dataResource MovieResource
  // Decode the incoming Movie json
  err := json.NewDecoder(request.Body).Decode(&dataResource);

  if err != nil {
    common.DisplayAppError(responseWriter, err, "Invalid Movie data", 500)

    return
  }

  movie := &dataResource.Data
  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("movies")
  // Insert a movie document
  repo := &data.MovieRepository{c}
  repo.Create(movie)
  j, err := json.Marshal(dataResource)

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}

// Handler for HTTP Get - "/movies/{id}"
// Get movie by id
func GetMovieById(responseWriter http.ResponseWriter, request *http.Request) {
  // Get id from incoming url
  vars := mux.Vars(request)
  id := vars["id"]

  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("movies")
  repo := &data.MovieRepository{c}

  // Get Movie by id
  movie, err := repo.GetById(id)
  if err != nil {
    if err == mgo.ErrNotFound {
      responseWriter.WriteHeader(http.StatusOK)

      return
    } else {
      common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

      return
    }
  }

  j, err := json.Marshal(MovieResource{Data:movie})

  if err != nil {
    common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

    return
  }

  responseWriter.Header().Set("Content-Type", "application/json")
  responseWriter.WriteHeader(http.StatusOK)
  responseWriter.Write(j)
}

// Handler for HTTP Delete - "/movies/{id}"
// Delete movie by id
func DeleteMovie(responseWriter http.ResponseWriter, request *http.Request) {
  // Get id from incoming url
  vars := mux.Vars(request)
  id := vars["id"]

  // Create new context
  context := NewContext()
  defer context.Close()

  c := context.DbCollection("movies")
  repo := &data.MovieRepository{c}
  err := repo.Delete(id)

  if err != nil {
    if err == mgo.ErrNotFound {
      responseWriter.WriteHeader(http.StatusNotFound)

      return
    } else {
      common.DisplayAppError(responseWriter, err, "An unexpected error has occurred", 500)

      return
    }
  }

  responseWriter.WriteHeader(http.StatusNoContent)
}
