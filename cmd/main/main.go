package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.open-meteo.com/v1/forecast?latitude=47.4984&longitude=19.0404&hourly=temperature_2m&timezone=Europe%2FBerlin&forecast_days=1"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
