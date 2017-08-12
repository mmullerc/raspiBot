package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Router struct {
	Mac      string
}

const (
	ROUTERS_COLLECTION = "routers"
)

//High order function
//Handles session.close when done
func queryWithRoutersCollection(fn func(*mgo.Collection) error) error {
	session := getMongoSession()
	defer session.Close()
	c := session.DB(DATABASE_NAME).C(ROUTERS_COLLECTION) // This can be parameterized later so it's more dynamic
	return fn(c)
}

func RouterExists(mac string) (exists bool, searchErr string) {
	searchErr = ""
	searchResults := Router{}
	exists = false

	//Query to run
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"mac": mac}).One(&searchResults)
		if searchResults.Mac == mac {
			exists = true
		}

		return fn
	}
	//calls High order function with query
	search := func() error {
		return queryWithRoutersCollection(query)
	}

	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return
}

//Inserts a new router in the database
func InsertRouter(mac string) string {
	searchErr := ""
	query := func(c *mgo.Collection) error {
		fn := c.Insert(&Router{mac})
		return fn
	}

	search := func() error {
		return queryWithRoutersCollection(query)
	}

	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return searchErr
}
