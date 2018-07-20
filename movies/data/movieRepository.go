package data

import (
  "gopkg.in/mgo.v2"
  "github.com/virtyaluk/go-simple-microservices-app/movies/models"
  "gopkg.in/mgo.v2/bson"
  "time"
)

type MovieRepository struct {
  C *mgo.Collection
}

func (r *MovieRepository) GetAll() []models.Movie {
  var movies []models.Movie
  iter := r.C.Find(nil).Iter()
  result := models.Movie{}

  for iter.Next(&result) {
    movies = append(movies, result)
  }

  return movies
}

func (r *MovieRepository) Create(movie *models.Movie) error {
  obj_id := bson.NewObjectId()
  movie.Id = obj_id
  movie.CreatedOn = time.Now()
  err := r.C.Insert(&movie)

  return err
}

func (r *MovieRepository) GetById(id string) (movie models.Movie, err error) {
  err = r.C.FindId(bson.ObjectIdHex(id)).One(&movie)

  return
}

func (r *MovieRepository) Delete(id string) error {
  err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

  return err
}
