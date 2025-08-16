package main

import (
	"flag"
	"fmt"

	"github.com/viacheslav-korobeynikov/weather/geo"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	//format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)

}
