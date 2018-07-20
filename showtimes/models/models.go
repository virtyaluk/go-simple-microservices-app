package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
)

type ShowTime struct {
  Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
  Date      string        `json:"date"`
  CreatedOn time.Time     `json:"createdon,omitempty"`
  Movies    []string      `json:"movies"`
}
