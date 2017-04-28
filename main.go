package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strings"
	_ "tinyUrl/routers"
)

func init() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
}

func main() {
	beego.Run()
}
