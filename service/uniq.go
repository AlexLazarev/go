package main

import (
	ai "github.com/night-codes/mgo-ai"
	"github.com/night-codes/tokay"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

//Worked with index that u can see in main function
func	create1(c *tokay.Context) {
	user := user{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	user.Id = ai.Next("users")
	if err := db.C("users").Insert(user); err != nil {
		c.JSON(400, obj{"err": "User already exists"})
		return
	}
	cache.add(&user)
	c.JSON(200, obj{"ok": "User has been created!"})
}


func	create2(c *tokay.Context) {
	user := user{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	uniq := obj{"login": strings.ToLower(user.Login)}
	if err := db.C("uniq").Find(uniq).One(nil); err == nil {
		c.JSON(400, obj{"err": "User already exist"})
		return
	}
	if err := db.C("uniq").Insert(uniq); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	user.Id = ai.Next("users")
	if err := db.C("users").Insert(user); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	cache.add(&user)
	c.JSON(200, obj{"ok": "User has been created!"})
}

func	create3(c *tokay.Context) {
	user := user{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	findParam := bson.M{"login": bson.RegEx{Pattern: user.Login, Options: "i"}}
	if err := db.C("users").Find(findParam).One(nil); err == nil {
		c.JSON(400, obj{"err": "User already exist"})
		return
	}
	user.Id = ai.Next("users")
	if err := db.C("users").Insert(user); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	cache.add(&user)
	c.JSON(200, obj{"ok": "User has been created!"})
}


func	create4(c *tokay.Context) {
	user := user{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	user.LowerLogin = strings.ToLower(user.Login)
	if err := db.C("users").Find(obj{"lower_login": user.LowerLogin}).One(nil); err == nil {
		c.JSON(400, obj{"err": "User already exist"})
		return
	}
	user.Id = ai.Next("users")
	if err := db.C("users").Insert(user); err != nil {
		c.JSON(400, obj{"err": err.Error()})
		return
	}
	cache.add(&user)
	c.JSON(200, obj{"ok": "User has been created!"})
}
