package controllers

import "github.com/virtyaluk/go-simple-microservices-app/movies/models"

type (
  // For Get - /movies
  MoviesResource struct {
    Data []models.Movie `json:"data"`
  }

  // For Post/Put - /movies
  MovieResource struct {
    Data models.Movie `json:"data"`
  }
)
