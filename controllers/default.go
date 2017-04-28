package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"tinyUrl/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["URL"] = ""
	c.Data["ShortURL"] = ""
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	url := c.GetString("url")
	shortUrl := c.GetString("short_url")
	c.TplName = "index.tpl"
	if url != "" {
		c.Data["URL"] = url
		c.Data["ShortURL"] = models.UrlEncode(url)
		return
	} else if shortUrl != "" {
		urls := strings.Split(shortUrl, "/")
		url := models.UrlDecode(urls[len(urls)-1])
		c.Data["URL"] = url
		c.Data["ShortURL"] = shortUrl
		return
	} else {
		c.Data["URL"] = ""
		c.Data["ShortURL"] = ""
		return
	}
}
