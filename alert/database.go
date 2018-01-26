package alert

import (
	"errors"
	"time"

	"github.com/davecgh/go-spew/spew"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	host = "mongodb://localhost:27017"
	// host            = "mongodb://adsethost.rebel9.io"
	alertCollection = "alerts"
	dbName          = "rebel"
)

// SendAlertAll Sends alert to all
func SendAlertAll(alert Alert) error {
	session, err := mgo.Dial(host)
	if err != nil {
		sendAlertAllNoStore(Alert{})
	}

	c := session.DB(dbName).C(alertCollection)
	defer session.Close()
	alert.DateCreated = time.Now()
	info, err := c.Upsert(&alert, &alert)
	if err != nil {
		return err
	}
	if info.UpsertedId == nil {
		return errors.New("Already exists")
	}
	a, err := getAlert(info.UpsertedId.(bson.ObjectId))
	if err != nil {
		return err
	}
	spew.Dump(a)
	//Alert Function

	return nil
}

func sendAlertAllNoStore(alert Alert) {
	//Alert function
}

func getAlerts() ([]Alert, error) {
	return nil, nil
}

func getAlert(id bson.ObjectId) (Alert, error) {
	// d1 := []byte("<html><p style='color:red'>hello</p></html>")
	session, err := mgo.Dial(host)
	if err != nil {
		SendAlertAll(Alert{})
	}

	c := session.DB(dbName).C(alertCollection)
	defer session.Close()
	alert := Alert{}
	err = c.Find(bson.M{"_id": id}).One(&alert)
	if err != nil {
		return Alert{}, err
	}
	return alert, nil
}
