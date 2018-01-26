package adsets

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// AdSet struct
// Overview: Used to store User Adsets
type AdSet struct {
	AdSetID     bson.ObjectId `bson:"_id,omitempty"`
	Link        string        `json:"Link,omitempty" bson:"Link,omitempty"`
	Status      string        `json:"Status,omitempty" bson:"Status,omitempty"`
	Type        string        `json:"Type,omitempty" bson:"Type,omitempty"`
	AdStats     []AdStat      `json:"AdStats,omitempty" bson:"AdStats,omitempty"`
	DateCreated time.Time     `json:"DateCreated,omitempty" bson:"DateCreated,omitempty"`
}

// AdStat struct
// Overview: Used to store Adsets statistics
type AdStat struct {
	Type        string    `json:"Type,omitempty" bson:"Type,omitempty"`
	Value       string    `json:"Value,omitempty" bson:"Value,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty" bson:"DateCreated,omitempty"`
}
