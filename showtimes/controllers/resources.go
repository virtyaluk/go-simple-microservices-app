package controllers

import "github.com/virtyaluk/go-simple-microservices-app/showtimes/models"

type (
  // For Get - /showtimes
  ShowTimesResource struct {
    Data []models.ShowTime `json:"data"`
  }

  // For Post/Put - /showtimes
  ShowTimeResource struct {
    Data models.ShowTime `json:"data"`
  }
)
