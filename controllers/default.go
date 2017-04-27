package controllers

import (
	"fmt"
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
		c.Data["ShortURL"] = urlEncode(url)
		return
	} else if shortUrl != "" {
		url := urlDecode(shortUrl)
		c.Data["URL"] = url
		c.Data["ShortURL"] = shortUrl
		return
	} else {
		c.Data["URL"] = ""
		c.Data["ShortURL"] = ""
		return
	}
}

func base62encode(number int) string {
	if number == 0 {
		return "0"
	}
	var i int
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62 := []string{}
	for number != 0 {
		number, i = number/62, number%62
		base62 = append(base62, string(alphabet[i]))
	}
	fmt.Println(base62)
	return strings.Join(base62, "")
}

func urlEncode(url string) string {
	shortUrl, _ := redis.Client.Get("reverse-url:" + url).Result()
	if shortUrl != "" {
		return shortUrl
	}
	shortIdInt, _ := redis.Client.Incr("last-url-id").Result()
	shortId := base62encode(int(shortIdInt))
	redis.Client.Set("url-target:"+shortId, url, 0)
	redis.Client.Set("reverse-url:"+url, shortId, 0)
	return shortId
}

func urlDecode(shortId string) string {
	url, _ := redis.Client.Get("url-target:" + shortId).Result()
	if url != "" {
		return url
	} else {
		return "Not Found"
	}
}
