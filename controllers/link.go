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
	url := models.UrlDecode(string(shortId[0]))
	fmt.Println(url)
	c.Redirect(url, 302)
}
