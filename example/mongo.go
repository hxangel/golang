package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProxyServerInfo struct {
	Host    string
	Rate    float64 //network speed
	Region  string  //region
	Type    int     //1 http 2 https 3 socket
	Status  bool    //region
	Level   uint8   //0 transparent 1 low 2 high
	Created int64   //timestamp created time
	Checked int64   //timestamp  checked time
}
type Proxy struct {
	List []*ProxyServerInfo
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("proxy").C("base")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

