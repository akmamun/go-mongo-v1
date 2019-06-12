package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Completed bool
}
