package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Weather struct {
	Properties struct {
		Periods []Period `json:"periods"`
	} `json:"properties"`
}

type Period struct {
	Temperature   int    `json:"temperature"`
	WindDirection string `json:"windDirection"`
	WindSpeed     string `json:"windSpeed"`
}

type Location struct {
	Loc string `json:"loc"`
}

type Office struct {
	Properties struct {
		GridId string `json:"gridId"`
		GridX int `json:"gridX"`
		GridY int `json:"gridY"`
	} `json:"properties"`
}

func fetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

func main() {
	office := flag.String("office", "", "NWS Office")
	lat := flag.String("lat", "", "Latitude")
	lon := flag.String("lon", "", "Longitude")
	gridX := flag.Int("x", -1, "Grid X")
	gridY := flag.Int("y", -1, "Grid Y")
	flag.Parse()

	var err error
	if *lat == "" || *lon == "" {
		var loc Location
		err = fetchJSON("https://ipinfo.io/geo", &loc)
		if err != nil {
			fmt.Println("Error fetching location:", err)
			return
		}
		coords := strings.Split(loc.Loc, ",")
		*lat, *lon = coords[0], coords[1]
	}

	if *office == "" {
		var off Office
		err = fetchJSON(fmt.Sprintf("https://api.weather.gov/points/%s,%s", *lat, *lon), &off)
		if err != nil {
			fmt.Println("Error fetching office:", err)
			return
		}
		*office = off.Properties.GridId
		*gridX = off.Properties.GridX
		*gridY = off.Properties.GridY
	}

	var weather Weather
	err = fetchJSON(fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast/hourly", *office, *gridX, *gridY), &weather)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}

	if len(weather.Properties.Periods) < 1 {
		fmt.Println("No weather data found")
		return
	}

	var high, low int
	high = weather.Properties.Periods[0].Temperature
	low = high

	for _, period := range weather.Properties.Periods {
		if period.Temperature > high {
			high = period.Temperature
		}
		if period.Temperature < low {
			low = period.Temperature
		}
	}

	now := weather.Properties.Periods[0]
	fmt.Printf("H %d°F L %d°F | ⛅ %d°F %s %s\n",
		high, low, now.Temperature, now.WindDirection, now.WindSpeed)
}

