package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os/exec"
)

type Information struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Userid string        `bson:"userid" json:"userid"`
	Size   string        `bson:"size" json:"size"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to REST API server!\n")
}

func createVolume(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "createVolume\n")
	decoder := json.NewDecoder(r.Body)
	var info Information
	err := decoder.Decode(&info)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %v \n", info.Userid, info.Size)

	// run script

	// to do list
	// 1) forbid duplicate insert
	// 2) error handling
	// 3) secure connection

	out, err := exec.Command("/bin/sh", "test.sh").Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(out))

	// db insertion
	session := dbconnection()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("volumes")

	if err != c.Insert(info) {
		panic(err)
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(r.Body)
}

func volumelist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//return volumelist in db
}

func deleteVolume(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//delete user volume
}

func resizeVolume(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//resize user volume
}

func dbconnection() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connect mongodb server to %v\n", session.LiveServers())
	return session.Clone()
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/volumes", createVolume)
	router.GET("/volumes", volumelist)
	router.DELETE("/volumes/:id", deleteVolume)
	router.PUT("/volumes/:id", resizeVolume)

	log.Fatal(http.ListenAndServe("localhost:1337", router))
}
