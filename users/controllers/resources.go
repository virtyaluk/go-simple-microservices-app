package controllers

import "github.com/virtyaluk/go-simple-microservices-app/users/models"

type (
  // For - /users
  UsersResource struct {
    Data []models.User `json:"data"`
  }

  // For Post/Put - /users
  UserResource struct {
    Data models.User `json:"data"`
  }
)
