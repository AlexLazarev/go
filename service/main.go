package main

import (
	ai "github.com/night-codes/mgo-ai"
	"github.com/night-codes/tokay"
	"gopkg.in/mgo.v2"
)

type obj map[string]interface{}

type user struct {
	Id 			uint64 `bson:"_id"`
	Login		string `form:"login" json:"login" bson:"login" valid:"required,min(3),max(40)"`
	Pass		string `form:"pass" json:"pass" bson:"pass" valid:"required,max(150)"`
	LowerLogin	string `bson:"lower_login"`
	Age			uint8 `form:"age" json:"age" bson:"age" valid:"required"`
}

type feedback struct {
	Id		uint64 `form:"id" bson:"_id"`
	Name    string `form:"name" bson:"name" valid:"required,min(3),max(40)"`
	Title   string `form:"title" bson:"title" valid:"required,max(150)"`
	Message string `form:"message" bson:"message" valid:"required"`
	Set 	bool   `bson:"set"`
	Date    string `bson:"time"`
	UserId  uint64 `bson:"userId"`
}

var (
	db *mgo.Database
)

func main() {
	session, err := mgo.Dial(":27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	app := tokay.New()
	db = session.DB("test2")
	ai.Connect(db.C("counters"))
	// for create1
	index := mgo.Index{
		Key: []string{"login"},
		Unique: true,
		Collation: &mgo.Collation{Locale: "en", Strength:1},
	}
	if err := db.C("users").EnsureIndex(index); err != nil {
		panic(err)
	}
	//
	go cache.worker_clear()
	go cache.worker_save()
	api := app.Group("/api")

	controllerUser(api.Group("/user"))

	panic(app.Run(":8082"))
}
