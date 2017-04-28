package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"tinyUrl/models"
)

type LinkController struct {
	beego.Controller
}

func (c *LinkController) Get() {
	shortUrl := c.GetString(":link")
	urls := strings.Split(shortUrl, "/")
	fmt.Println(shortUrl)
	url := models.UrlDecode(urls[len(urls)-1])
	fmt.Println(url)
	c.Redirect(url, 302)
}
