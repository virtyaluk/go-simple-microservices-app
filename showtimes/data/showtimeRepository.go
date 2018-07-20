package data

import (
  "gopkg.in/mgo.v2"
  "github.com/virtyaluk/go-simple-microservices-app/showtimes/models"
  "gopkg.in/mgo.v2/bson"
  "time"
)

type ShowtimeRepository struct {
  C *mgo.Collection
}

func (r *ShowtimeRepository) Create(showtime *models.ShowTime) error {
  obj_id := bson.NewObjectId()
  showtime.Id = obj_id
  showtime.CreatedOn = time.Now()
  err := r.C.Insert(&showtime)

  return err
}

func (r *ShowtimeRepository) GetAll() []models.ShowTime {
  var showtimes []models.ShowTime
  iter := r.C.Find(nil).Iter()
  result := models.ShowTime{}

  for iter.Next(&result) {
    showtimes = append(showtimes, result)
  }

  return showtimes
}

func (r *ShowtimeRepository) GetByDate(date string) (showtime models.ShowTime, err error) {
  err = r.C.Find(bson.M{"date": date}).One(&showtime)

  return
}

func (r *ShowtimeRepository) Delete(id string) error {
  err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

  return err
}
