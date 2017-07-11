package db

import (
  "fmt"
  "log"
  "net/http"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type State struct {
  Component string
  State string
  Direction string
  Speed string 
}

func GetStateByComponent(w http.ResponseWriter, component string) {
  session, err := mgo.Dial("10.28.6.16")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  c := session.DB("test").C("state")
  result := State{}
  err = c.Find(bson.M{"component": component}).One(&result)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Fprint(w, result.Component, ": ", result.State, ",", result.Direction, ",", result.Speed, "\n")
}

// func InsertState(component string, state string, direction string, speed string) {
//   session, err := mgo.Dial("10.28.6.16")
//   if err != nil {
//     panic(err)
//   }
//   defer session.Close()

//   c := session.DB("test").C("state")
//   err = c.Insert(&State{component, state, direction, speed})
//   if err != nil {
//     log.Fatal(err)
//   }
// }
