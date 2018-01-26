package adsets

import (
	"errors"
	"time"

	"github.com/dickmanben/rebel-api/alert"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	host = "mongodb://localhost:27017"
	// host            = "mongodb://adsethost.rebel9.io"
	adsetcollection = "adsets"
	dbName          = "rebel"
)

func getPassThru(id bson.ObjectId) ([]byte, error) {
	adset, err := getAdset(id)
	if err != nil {
		return nil, err
	}

	content := []byte("<html><head><script src='/js/track.js'></script></head><body><p>If you are not redirected automatically please <a id='redirect' href='" + adset.Link + "'>Click here</a></p></body><script type='text/javascript' defer>console.log('done');document.getElementById('redirect').click()</script></html>")
	return content, nil
}

func createAdset(a AdSet) (AdSet, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		alert.SendAlertAll(alert.Alert{})
		return AdSet{}, err
	}

	c := session.DB(dbName).C(adsetcollection)
	defer session.Close()
	a.DateCreated = time.Now()
	info, err := c.Upsert(&a, &a)
	if err != nil {
		return AdSet{}, err
	}
	if info.UpsertedId == nil {
		return AdSet{}, errors.New("Already exists")
	}
	adset, err := getAdset(info.UpsertedId.(bson.ObjectId))
	if err != nil {
		return AdSet{}, err
	}
	return adset, nil
}

func updateAdset(a AdSet, id bson.ObjectId) (AdSet, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		alert.SendAlertAll(alert.Alert{})
		return AdSet{}, err
	}

	c := session.DB(dbName).C(adsetcollection)
	defer session.Close()
	query := bson.M{"_id": id}
	update := bson.M{"$set": &a}

	err = c.Update(query, update)
	if err != nil {
		return AdSet{}, err
	}

	adset, err := getAdset(id)
	if err != nil {
		return AdSet{}, err
	}
	return adset, nil
}

func getAdset(id bson.ObjectId) (AdSet, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		alert.SendAlertAll(alert.Alert{})
	}

	c := session.DB(dbName).C(adsetcollection)
	defer session.Close()
	adset := AdSet{}
	err = c.Find(bson.M{"_id": id}).One(&adset)
	if err != nil {
		return AdSet{}, err
	}
	return adset, nil
}

func getAdsets() ([]AdSet, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		alert.SendAlertAll(alert.Alert{})
	}

	c := session.DB(dbName).C(adsetcollection)
	defer session.Close()
	adsets := []AdSet{}
	err = c.Find(nil).All(&adsets)
	if err != nil {
		return nil, err
	}
	return adsets, nil
}
