package main

import (
	"github.com/night-codes/tokay"
)
func controllerUser(gr *tokay.RouterGroup) {
	gr.POST("/create", create1)

	gr.GET("/findUser", func(c *tokay.Context) {
		login := c.Query("login")
		user := user{}
		if err := cache.get(login, &user); err != nil {
			c.JSON(400, obj{"err": "Login not found"})
		} else {
			c.JSON(200, user)
		}
	})

	gr.GET("/updateAge", func(c *tokay.Context) {
		id := uint64(c.QueryUint("id"))
		age := uint8(c.QueryUint("age"))
		if err := cache.update(id, age); err != nil {
			c.JSON(400, obj{"err": err})
		} else {
			c.JSON(200, obj{"ok": "User has been updated!"})
		}
	})
}