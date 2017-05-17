package models

import (
	"os"
	"fmt"
	"strings"
)

var BASE_URL string

func init() {
	if baseUrl := os.Getenv("BASE_URL"); baseUrl != "" {
		BASE_URL = os.Getenv("BASE_URL")
	} else {
		BASE_URL = "t.instask.me"
	}
}

func Base62Encode(number int) string {
	if number == 0 {
		return "000000"
	}
	var i int
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62 := []string{}
	for number != 0 {
		number, i = number/62, number%62
		base62 = append(base62, string(alphabet[i]))
	}
	r := strings.Join(base62, "")
	return fmt.Sprintf("%06s", r)
}

func UrlEncode(url string) string {
	if sid, _ := Client.Get("reverse-url:" + url).Result(); sid != "" {
		return BASE_URL + "/link/" + sid
	}
	shortIdInt, _ := Client.Incr("last-url-id").Result()
	shortId := Base62Encode(int(shortIdInt))
	fmt.Println(shortId)
	Client.Set("url-target:" + shortId, url, 0)
	Client.Set("reverse-url:" + url, shortId, 0)
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
