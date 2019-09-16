package main

import (
	ai "github.com/night-codes/mgo-ai"
	"github.com/night-codes/tokay"
	"github.com/night-codes/types"
	"log"
	"time"
)

func ControllerUser(app *tokay.Engine) {

	app.GET("/registration", func(c *tokay.Context) {
		c.HTML(200, "registration", nil)
	})

	app.GET("/authorization", func(c *tokay.Context) {
		c.HTML(200, "authorization", nil)
	})

	app.GET("/logout", func(c *tokay.Context) {
		c.RemoveCookie("login")
		c.RemoveCookie("pass")
		c.Redirect(302, "/authorization")
	})

	app.GET("/content", func(c *tokay.Context) {
		feedback := []feedback{}
		log.Println(c.Get("id"))
		if err:= db.C("feedbacks").Find(obj{"set" : true}).Sort("-time").All(&feedback); err != nil {
			c.String(400, "Some trouble")
		} else {
			c.HTML(200, "content", obj{"feedbacks": feedback})
		}
	})

	app.GET("/", func(c *tokay.Context) {
		if (c.Cookie("login") == "") {
			c.Redirect(302, "/authorization")
		} else {
			c.Redirect(302, "/u/personal")
		}
	})

	app.POST("/authorization", func(c *tokay.Context) {
		orig := user{}
		user := user{}
		ret := obj{}
		ex := false
		if user.Login, ex = c.PostFormEx("login"); ex == false {
			ret["err"] = "Oops, an error: login"
		} else if user.Pass, ex = c.PostFormEx("pass"); ex == false {
			ret["err"] = "Oops, an error: pass"
		} else if err := api.getUser(user.Login, &orig); err != nil {
			ret["err"] = "Login is not registered"
		} else if (!CheckPasswordHash(user.Pass, orig.Pass)) {
			ret["err"] = "Wrong password"
		} else {
			ret["ok"] = "Authorization success"
			c.SetCookie("login", orig.Login, "", "", false, false)
			c.SetCookie("pass", orig.Pass, "", "", false, false)
			//c.SetCookie("id", fmt.Sprint(orig.Id), "", "", false, false)
			//c.Set("id", orig.Id)
			//log.Println("ID", types.Uint64(c.Get("id")), orig.Id)
		}
		c.HTML(200, "authorization",  ret)
	})

	app.POST("/registration", func(c *tokay.Context) {
		user := user{}
		ret := obj{}
		if err := c.Bind(&user); err != nil {
			ret["err"] = "Oops, an error: " + err.Error()
		} else if user.Pass, err = HashPassword(user.Pass); err != nil {
			ret["err"] = err
		} else {
			if err = api.createUser(user); err != nil{
				ret["err"] = err
			} else {
				ret["ok"] = "Successful!"
			}
		}
		c.HTML(200, "registration", ret)
	})


	gr := app.Group("/u", func(c *tokay.Context){
		user := user{}
		err := api.getUser(c.Cookie("login"), &user)
		if  err != nil || (c.Cookie("pass") != user.Pass) {
			c.Abort()
			c.Redirect(302, "/authorization")
		} else {
			c.Set("id", user.Id)
		}
	})

	gr.GET("/", func(c *tokay.Context) {
		c.Redirect(302, "/u/personal")
	})


	gr.GET("/feedback", func(c *tokay.Context) {
		c.HTML(200, "feedback", nil)
	})

	gr.GET("/personal/update", func(c *tokay.Context) {
		c.HTML(200, "update", nil)
	})

	gr.POST("/personal/update", func(c *tokay.Context) {
		ret := obj{}
		age := types.Uint64(c.PostForm("age"))
		if err := api.updateAge(types.Uint64(c.Get("id")), age); err != nil{
			ret["err"] = "WTF " + err.Error()
		} else {
			ret["ok"] = "Successful!"
		}
		c.HTML(200, "update", ret)
	})

	gr.GET("/personal", func(c *tokay.Context) {
		user := user{}
		fb := []feedback{}
		if err := api.getUser(c.Cookie("login"), &user); err != nil {
			c.String(400, err.Error())
		} else if err = db.C("feedbacks").Find(obj{"userId": user.Id}).All(&fb); err != nil {
			c.String(400, err.Error())
		} else {
			c.HTML(200, "personal", obj{"u": user, "fb": fb})
		}
	})

	gr.POST("/feedback", func(c *tokay.Context) {
		fb := feedback{}
		user := user{}
		ret := obj{}
		if err := c.Bind(&fb); err != nil {
			ret["err"] = "Oops, an error: " + err.Error()
		} else if err = api.getUser(c.Cookie("login"), &user); err != nil {
			ret["err"] = err
		} else {
			fb.Id = ai.Next("feedbacks")
			fb.Name = c.Cookie("login")
			fb.UserId = user.Id
			fb.Set = false
			t := time.Now()
			fb.Date = t.Format("20060102150405")
			if err = db.C("feedbacks").Insert(fb); err != nil {
				ret["err"] = "Unexpected error. Come back to us later."
			} else {
				ret["ok"] = "Thanks for your feedback!"
			}
		}
		c.HTML(200, "feedback", ret)
	})
}
