package mongo

import "gopkg.in/mgo.v2"

var (
	// Db exposes a global database connection
	Db *Mongo
)

// Mongo is a service wrapper to share a connection to a mongodb
type Mongo struct {
	Session *mgo.Session
}

// Connect to a mongodb
func (mongo *Mongo) Connect(host string) error {
	session, err := mgo.Dial(host)
	if err != nil {
		return err
	}

	mongo.Session = session
	return nil
}

// Collection returns a document collection
func (mongo *Mongo) Collection(db string, name string) *mgo.Collection {
	return mongo.Session.DB(db).C(name)
}

func init() {
	Db = &Mongo{}
}
