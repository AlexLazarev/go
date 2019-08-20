package main

import (
	"fmt"
	ai "github.com/night-codes/mgo-ai"
	"github.com/night-codes/tokay"
	"gopkg.in/mgo.v2"
	"log"
)

func controllerUser(gr tokay.RouterGroup, db mgo.Database) {
	gr.GET("/create", func(c *tokay.Context) {
		user := user{}
		ret := obj{}
		if err := c.Bind(&user); err != nil {
			ret["err"] = "Oops, an error: " + err.Error()
		} else {
			user.Id = ai.Next("Users")
			if err := db.C("Users").Insert(user); err != nil {
				ret["err"] = "Unexpected error. Come back to us later."
			} else {
				ret["ok"] = fmt.Sprint("User ", user.Id, " has been created!")
			}
		}
		c.JSON(200, ret)
	})

	gr.GET("/delete", func(c *tokay.Context) {
		id := uint64(c.QueryUint("id"))
		log.Println(id)
		ret := obj{}
		if err := db.C("Users").Remove(obj{"_id": id}); err != nil {
			ret["err"] = fmt.Sprint("User with id ", id, " was not found")
		} else {
			ret["ok"] = fmt.Sprint("User ", id, " has been deleted!")
		}
		c.JSON(200, ret)
	})

	gr.GET("/findLogin", func(c *tokay.Context) {
		login := c.Query("login")
		ret := obj{}
		if err := db.C("Users").Find(obj{"login": login}).Select(obj{"_id": true}).One(&ret); err != nil {
			ret["err"] = fmt.Sprint("User ", login, " is not found")
		}
		c.JSON(200, ret)
	})

	gr.GET("/get", func(c *tokay.Context) {
		id := uint64(c.QueryUint("id"))
		ret := obj{}
		if err := db.C("Users").FindId(id).One(&ret); err != nil {
			ret["err"] = fmt.Sprint("User ", id, " is not found")
		}
		c.JSON(200, ret)
	})
}