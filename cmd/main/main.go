package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Forecast struct {
	Hourly struct {
		Time        []string  `json:"time"`
		Temperature []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

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

	var data Forecast

	error := json.Unmarshal(body, &data)

	if error != nil {
		panic(error)
	}

	for i, value := range data.Hourly.Time {
		t, _ := time.Parse("2006-01-02T15:04", value)

		if t.Before(time.Now()) {
			continue
		}

		line := fmt.Sprintf("%s: %.0f°", t.Format("15:04"), data.Hourly.Temperature[i])

		fmt.Println(line)
	}

}
