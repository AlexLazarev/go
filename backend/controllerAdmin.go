package main

import (
	"fmt"
	"github.com/night-codes/tokay"
)

func ControllerAdmin(app *tokay.Engine) {

	rg := app.Group("/admin", tokay.BasicAuth("admin", "admin"))

	rg.GET("", func(c *tokay.Context) {
		c.Redirect(302, rg.Path() + "/feedbacks")
	})

	rg.GET("/feedbacks", func(c *tokay.Context) {
		feedback := []feedback{}
		if err:= db.C("feedbacks").Find(nil).All(&feedback); err != nil {
			c.String(400, "Some trouble")
		} else {
			c.HTML(200, "feedbacks", obj{"feedbacks": feedback})
		}
	})
	rg.GET("/feedbacks/delete/<id>", func(c *tokay.Context) {
		id := uint64(c.ParamUint("id"))
		if err := db.C("feedbacks").RemoveId(id); err != nil {
			c.String(400, fmt.Sprint("Group with id ", id, " was not found"))
		} else {
			c.Redirect(302, "/admin/feedbacks")
		}
	})

	rg.GET("/feedbacks/set/<id>", func(c *tokay.Context) {
		id := uint64(c.ParamUint("id"))
		if err := db.C("feedbacks").UpdateId(id, obj{"$set": obj{"set": true}}); err != nil {
			c.String(400, fmt.Sprint("Group with id ", id, " was not found"))
		} else {
			c.Redirect(302, "/admin/feedbacks")
		}
	})

	rg.GET("/feedbacks/unset/<id>", func(c *tokay.Context) {
		id := uint64(c.ParamUint("id"))
		if err := db.C("feedbacks").UpdateId(id, obj{"$set": obj{"set": false}}); err != nil {
			c.String(400, fmt.Sprint("Group with id ", id, " was not found"))
		} else {
			c.Redirect(302, "/admin/feedbacks")
		}
	})

/*	rg.POST("/feedbacks", func(c *tokay.Context) {
		log.Println("HEH")
		id := id{}
		ret := obj{}
		if err := c.Bind(&id); err != nil {
			ret["err"] = "Oops, an error: " + err.Error()
		} else if err = db.C("feedbacks").RemoveId(id.Id); err != nil {
			ret["err"] = fmt.Sprint("Group with id ", id.Id, " was not found")
		} else {
			ret["ok"] = fmt.Sprint("Group ", id.Id, " has been deleted!")
		}
		log.Println(id.Id)
		c.HTML(200, "feedbacks", ret)
	})*/



}