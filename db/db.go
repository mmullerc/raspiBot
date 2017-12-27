package db

import (
	"log"
	"raspibot/utilities"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Database session
var (
	mgoSession *mgo.Session
)

type Component struct {
	Name      string
	State     string
	Direction string
	Speed     string
}

const (
	MONGO_CONN_STR        = "10.28.6.16"
	DATABASE_NAME         = "raspiBot"
	COMPONENTS_COLLECTION = "components"
	USERS_COLECTION 	  =	"users"
)

//Retruns a NEW session, but reuses the same socket as the original session
func getMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(MONGO_CONN_STR)
		if err != nil {
			log.Fatal("Failed to start MongoDB session")
		}
	}
	return mgoSession.Clone()
}

//High order function
//Handles session.close when done
func queryWithComponentsCollection(fn func(*mgo.Collection) error) error {
	session := getMongoSession()
	defer session.Close()
	c := session.DB(DATABASE_NAME).C(COMPONENTS_COLLECTION) // This can be parameterized later so it's more dynamic
	return fn(c)
}

//High order function
//Handles session.close when done
func queryWithUsersCollection(fn func(*mgo.Collection) error) error {
	session := getMongoSession()
	defer session.Close()
	c := session.DB(DATABASE_NAME).C(USERS_COLECTION) // This can be parameterized later so it's more dynamic
	return fn(c)
}

//Retrieves a component by it's name
func GetStateByComponent(name string) (searchResults Component, searchErr string) {
	searchErr = ""
	searchResults = Component{}

	//Query to run
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"name": name}).One(&searchResults)
		return fn
	}
	//calls High order function with query
	search := func() error {
		return queryWithComponentsCollection(query)
	}

	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return
}

//Inserts a new component in the database
func insertState(component string, state string, direction string, speed string) string {
	searchErr := ""
	query := func(c *mgo.Collection) error {
		fn := c.Insert(&Component{component, state, direction, speed})
		return fn
	}

	search := func() error {
		return queryWithComponentsCollection(query)
	}

	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return searchErr
}

//Updates the components state
func updateState(state string) string {
	searchErr := ""
	selector := bson.M{"name": "car"}
	update := bson.M{"$set": bson.M{"state": state}}

	query := func(c *mgo.Collection) error {
		fn := c.Update(selector, update)
		return fn
	}

	search := func() error {
		return queryWithComponentsCollection(query)
	}

	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return searchErr
}

//Updates the components direction
func UpdateDirection(direction string) string {
	searchErr := ""
	selector := bson.M{"name": "car"}
	update := bson.M{"$set": bson.M{"direction": direction}}

	query := func(c *mgo.Collection) error {
		fn := c.Update(selector, update)
		return fn
	}

	search := func() error {
		return queryWithComponentsCollection(query)
	}

	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return searchErr
}


//Starts the state, if empty it creates one
func StartCar() {
	car, err := GetStateByComponent("car")
	utilities.CheckForStringErr(err)

	if (Component{}) == car {
		err = insertState("car", "on", "", "")
		utilities.CheckForStringErr(err)
	} else {
		err = updateState("on")
		utilities.CheckForStringErr(err)
	}
}

//Updates the state, if empty creates one
func StopCar() {
	car, err := GetStateByComponent("car")
	utilities.CheckForStringErr(err)

	if (Component{}) == car {
		err = insertState("car", "off", "", "")
		utilities.CheckForStringErr(err)
	} else {
		err = updateState("off")
		utilities.CheckForStringErr(err)
	}
}
