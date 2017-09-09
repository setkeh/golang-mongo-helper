package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DoQuery Run Mongo Query
func DoQuery(session *mgo.Session, c Config, q Query) (map[string]interface{}, error) {
	conn := session.DB(c.Database).C(q.Collection)

	//result := make([]bson.M, 50)
	result := make(map[string]interface{})

	conn.Find(bson.M{q.QueryKey: q.QueryValue}).One(&result)

	return result, nil
}
