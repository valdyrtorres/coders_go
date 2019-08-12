package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	ID   bson.ObjectId `bson:"_id,onitempty" json:"id"`
	Name string        `bson:"name" json:"name"`
	Desc string        `bson:"desc" json:"desc"`
}

var Tasks = new(tasks)

type tasks struct{}
