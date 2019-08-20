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
	Age			string `form:"age" json:"age" bson:"age" valid:"required"`
	GroupId		[]uint64 `json:"groupId" bson:"groupId"`
}

type group struct {
	Id 			uint64 `bson:"_id"`
	Title   	string `form:"title" json:"title" bson:"title" valid:"required,max(150)"`
}

func main() {
	session, err := mgo.Dial(":27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	app := tokay.New()

	ai.Connect(session.DB("mydb").C("counters"))

	api := app.Group("/api")

	controllerUser(*api.Group("/user"), *session.DB("mydb"))
	controllerGroup(*api.Group("/group"), *session.DB("mydb"))

	app.Run(":8080")
}
