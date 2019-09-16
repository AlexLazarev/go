package main

import (
	ai "github.com/night-codes/mgo-ai"
	"github.com/night-codes/tokay"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

type obj map[string]interface{}

type feedback struct {
	Id		uint64 `form:"id" bson:"_id"`
	Name    string `form:"name" bson:"name"`
	Title   string `form:"title" bson:"title" valid:"required,max(150)"`
	Message string `form:"message" bson:"message" valid:"required"`
	Set 	bool   `bson:"set"`
	Date    string `bson:"time"`
	UserId  uint64 `bson:"userId"`
}

type user struct {
	Id 			uint64 `json:"id" bson:"_id"`
	Login		string `form:"login" json:"login" bson:"login" valid:"required,min(3),max(40)"`
	Pass		string `form:"pass" json:"pass" bson:"pass" valid:"required,max(150)"`
	Age			uint8 `form:"age" json:"age" bson:"age" valid:"required, min(1), max(120)"`
}

type API  struct {
	address string
	user string
}

var (
	db *mgo.Database
)

func main() {
	session, err := mgo.Dial(":27017")
	if err != nil {
		panic(err)
	}
	app := tokay.New()
	app.Static("/files", "./files")

	db = session.DB("test2")
	ai.Connect(db.C("counters"))
	ControllerUser(app)
	ControllerAdmin(app)

	app.Run(":8080")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
