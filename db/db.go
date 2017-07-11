package db

import (
  "fmt"
  "log"
  "net/http"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Component struct {
  Name string
  State string
  Direction string
  Speed string 
}

func GetStateByComponent(w http.ResponseWriter, name string) {
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

  fmt.Fprint(w, result.Name, ": ", result.State, ",", result.Direction, ",", result.Speed, "\n")
}

// func InsertState(component string, state string, direction string, speed string) {
//   session, err := mgo.Dial("10.28.6.16")
//   if err != nil {
//     panic(err)
//   }
//   defer session.Close()

//   c := session.DB("raspiBot").C("state")
//   err = c.Insert(&Component{component, state, direction, speed})
//   if err != nil {
//     log.Fatal(err)
//   }
// }
