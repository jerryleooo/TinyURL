package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"tinyUrl/models"
)

type LinkController struct {
	beego.Controller
}

func (c *LinkController) Get() {
	shortId := c.GetString(":link")
	fmt.Println(shortId)
	url := models.UrlDecode(shortId)
	c.Redirect(url, 302)
}
