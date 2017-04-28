package models

import (
	"fmt"
	"strings"
)

var BASE_URL string

func init() {
	BASE_URL = "www.w2read.com:10080"
}

func Base62Encode(number int) string {
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

func UrlEncode(url string) string {
	shortUrl, _ := Client.Get("reverse-url:" + url).Result()
	if shortUrl != "" {
		return BASE_URL + "/link/" + shortUrl
	}
	shortIdInt, _ := Client.Incr("last-url-id").Result()
	shortId := Base62Encode(int(shortIdInt))
	Client.Set("url-target:"+shortId, url, 0)
	Client.Set("reverse-url:"+url, shortId, 0)
	return BASE_URL + "/link/" + shortId
}

func UrlDecode(shortId string) string {
	url, _ := Client.Get("url-target:" + shortId).Result()
	if url != "" {
		return url
	} else {
		return "Not Found"
	}
}
