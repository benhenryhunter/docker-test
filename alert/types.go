package alert

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Alert struct
// Overview: Used to structure Alerts sent from system
type Alert struct {
	AlertID     bson.ObjectId `bson:"_id,omitempty"`
	Severity    int           `json:"Severity,omitempty" bson:"Severity,omitempty"`
	Value       string        `json:"Value,omitempty" bson:"Value,omitempty"`
	Type        string        `json:"Type,omitempty" bson:"Type,omitempty"`
	DateCreated time.Time     `json:"DateCreated,omitempty" bson:"DateCreated,omitempty"`
}
