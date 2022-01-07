package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"key/models"
	"net/http"
)

func Weather(name string) {
	res, err := http.Get("https://restapi.amap.com/v3/weather/weatherInfo?city=" + name +
		"&key=e56a2520ea560138ea4da023cfa7fb91")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	resGet, resErr := ioutil.ReadAll(res.Body)
	if resErr != nil {
		fmt.Println("读取失败", resErr)
	}
	var data models.Weather
	err = json.Unmarshal([]byte(resGet), &data)
	if err != nil {
		fmt.Println("反序列化失败:", err)
	}
	fmt.Println(data)
}
