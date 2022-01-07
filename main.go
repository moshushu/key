package main

import (
	"flag"
	"fmt"
	"key/controllers"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "Please input city name")
	flag.Usage = func() {
		fmt.Println("帮助信息:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(name) == 0 {
		fmt.Println("请输入城市编号")
		return
	}
	controllers.Weather(name)
}
