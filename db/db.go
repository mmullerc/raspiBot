package db

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Component struct {
	Name      string
	State     string
	Direction string
	Speed     string
}

func GetStateByComponent(name string) Component {
	session, err := mgo.Dial("10.28.6.16")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("raspiBot").C("components")
	result := Component{}
	err = c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func openSession() *Collection {
  session, err := mgo.Dial("10.28.6.16")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  c := session.DB("raspiBot").C("components")
  return c
}

func InsertState(component string, state string, direction string, speed string) {
  c := openSession()
  err = c.Insert(&Component{component, state, direction, speed})
  if err != nil {
    log.Fatal(err)
  }
}

func UpdateState(component string, state string, direction string, speed string) {
  c := openSession()
  selector := bson.M{"Name":"car"}
  update := bson.M{"$inc": bson.M{"Name": component, "State": state, "Direction":direction, "Speed": speed}}

  err := c.Update(selector, update)
  if err != nil {
    panic(err)
  }

}

func StartCar() {
	car := GetStateByComponent("car")
	if car == nil {
		InsertState("car","on",nil,nil)
	} else {
		UpdateState("car","om", nil, nil)
	}
}

func StopCar() {
	car := GetStateByComponent("car")
	if car == nil {
		InsertState("car","off",nil,nil)
	} else {
		UpdateState("car","off", nil, nil)
	}
}
