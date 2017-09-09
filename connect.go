package mongo

import (
	"sync"

	"gopkg.in/mgo.v2"
)

var (
	waitGroup sync.WaitGroup
)

// Connect to Mongo Instance
func Connect(c Config) (*mgo.Session, error) {
	dialinfo := &mgo.DialInfo{
		Addrs:    c.Host,
		Database: c.Database,
		Username: c.Username,
		Password: c.Password,
	}

	session, err := mgo.DialWithInfo(dialinfo)
	if err != nil {
		return nil, err
	}

	return session, nil
}
