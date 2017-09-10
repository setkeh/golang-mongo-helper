package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DoReadQuery Read Data From Mongo
func DoReadQuery(session *mgo.Session, c Config, q Query) (map[string]interface{}, error) {
	conn := session.DB(c.Database).C(c.Collection)

	//result := make([]bson.M, 50)
	result := make(map[string]interface{})

	conn.Find(bson.M{q.QueryKey: q.QueryValue}).One(&result)

	return result, nil
}

//DoInsertQuery Inset a Document into Mongo
func DoInsertQuery(session *mgo.Session, c Config, q interface{}) error {
	conn := session.DB(c.Database).C(c.Collection)

	err := conn.Insert(q)
	if err != nil {
		return err
	}

	return nil
}
