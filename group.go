package main

import (
	"fmt"
	ai "github.com/night-codes/mgo-ai"
	"github.com/night-codes/tokay"
	"gopkg.in/mgo.v2"
	"log"
)

func controllerGroup(gr tokay.RouterGroup, db mgo.Database) {
	gr.GET("/create", func(c *tokay.Context) {
		group := group{}
		ret := obj{}
		if err := c.Bind(&group); err != nil {
			ret["err"] = "Oops, an error: " + err.Error()
		} else {
			log.Println(group)
			group.Id = ai.Next("Groups")
			if err := db.C("Groups").Insert(group); err != nil {
				ret["err"] = "Unexpected error. Come back to us later."
			} else {
				ret["id"] = fmt.Sprint("Group ", group.Id, " has been created!")
			}
		}
		c.JSON(200, ret)
	})

	gr.GET("/delete", func(c *tokay.Context) {
		id := uint64(c.QueryUint("id"))
		log.Println(id)
		ret := obj{}
		if err := db.C("Groups").Remove(obj{"_id": id}); err != nil {
			ret["err"] = fmt.Sprint("Group with id ", id, " was not found")
		} else {
			ret["ok"] = fmt.Sprint("Group ", id, " has been deleted!")
		}
		c.JSON(200, ret)
	})


	gr.GET("/addUser", func(c *tokay.Context) {
		id := uint64(c.QueryUint("id"))
		userId := uint64(c.QueryUint("userId"))
		ret := obj{}

		if err := db.C("Groups").FindId(id).One(nil); err != nil {
			ret["err"] = fmt.Sprint("Group ", id, " is not found")
		} else if err := db.C("Users").UpdateId(userId, obj{"$addToSet": obj{"groupId": id}}); err != nil {
			ret["err"] = fmt.Sprint("User ", userId, " is not found")
		} else {
			ret["ok"] = fmt.Sprint("User ", userId, " has been to group ", id)
		}
		c.JSON(200, ret)
	})


	gr.GET("/getUsers", func(c *tokay.Context) {
		id := uint64(c.QueryUint("id"))
		ret := []obj{}
		if err:= db.C("Groups").FindId(id).One(nil); err != nil {
			c.JSON(200, fmt.Sprint("Group ", id, " is not found"))
		} else if err:= db.C("Users").Find(obj{"groupId": id}).All(&ret); err != nil {
			c.JSON(200, "Some trouble")
		} else if len(ret) == 0 {
			c.JSON(200, fmt.Sprint("Group ", id, " doesn't have any subscribers"))
		} else {
			c.JSON(200, ret)
		}
	})

	gr.GET("/deleteUser", func(c *tokay.Context) {
		id := uint64(c.QueryUint("id"))
		userId := uint64(c.QueryUint("userId"))
		ret := obj{}
		if err:= db.C("Groups").FindId(id).One(nil); err != nil {
			c.JSON(200, fmt.Sprint("Group ", id, " is not found"))
		} else if err := db.C("Users").UpdateId(userId, obj{"$pull": obj{"groupId": id}}); err != nil {
			ret["err"] = fmt.Sprint("User ", userId, " is not found")
		} else {
			ret["ok"] = "true"
		}
		c.JSON(200, ret)
	})

	gr.GET("/plus", func(c *tokay.Context) {
		id1 := uint64(c.QueryUint("id1"))
		id2 := uint64(c.QueryUint("id2"))
		users := []user{}
		ret := obj{}

		if err := db.C("Groups").FindId(id1).One(nil); err != nil {
			ret["err"] = fmt.Sprint("Group ", id1, " is not found")
		} else if err = db.C("Groups").FindId(id2).One(nil); err != nil {
			ret["err"] = fmt.Sprint("Group ", id2, " is not found")
		} else if err := db.C("Users").Find(obj{"groupId": id2}).All(&users); err != nil {
			ret["err"] = fmt.Sprint("Group ", id2, " doesn't have any subscribers")
		} else {
			for i := 0; i < len(users); i++ {
				if err := db.C("Users").UpdateId(users[i].Id, obj{"$addToSet": obj{"groupId": id1}}); err != nil {
					ret["err"] = fmt.Sprintln("Some trouble with user ", users[i].Id)
				}
			}
			ret["ok"] = true
		}
		c.JSON(200, ret)
	})

	gr.GET("/minus", func(c *tokay.Context) {
		id1 := uint64(c.QueryUint("id1"))
		id2 := uint64(c.QueryUint("id2"))
		users := []user{}
		ret := obj{}

		if err := db.C("Groups").FindId(id1).One(nil); err != nil {
			ret["err"] = fmt.Sprint("Group ", id1, " is not found")
		} else if err = db.C("Groups").FindId(id2).One(nil); err != nil {
			ret["err"] = fmt.Sprint("Group ", id2, " is not found")
		} else if err := db.C("Users").Find(obj{"groupId": id2}).All(&users); err != nil {
			ret["err"] = fmt.Sprint("Group ", id2, " doesn't have any subscribers")
		} else {
			for i := 0; i < len(users); i++ {
				if err := db.C("Users").UpdateId(users[i].Id, obj{"$pull": obj{"groupId": id1}}); err != nil {
					ret["err"] = fmt.Sprintln("Some trouble with user ", users[i].Id)
				}
			}
			ret["ok"] = true
		}
		c.JSON(200, ret)
	})
}




