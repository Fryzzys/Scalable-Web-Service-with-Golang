package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Weather struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"status water"`
	WindStatus  string `json:"status wind"`
}

func main() {
	timer := time.NewTicker(15 * time.Second)
	defer timer.Stop()

	for range timer.C {
		id := rand.Intn(100) + 1

		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?id=%d", id)

		//fmt.Printf("Sending HTTP request to %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error sending HTTP request:", err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Printf("Error response status code: %d %s\n", resp.StatusCode, resp.Status)
			continue
		}

		var posts []Post
		err = json.NewDecoder(resp.Body).Decode(&posts)
		if err != nil {
			resp.Body.Close()
			fmt.Println("Error decoding response body:", err)
			continue
		}
		resp.Body.Close()

		var waterStatus, windStatus string
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		if water <= 5 {
			waterStatus = "aman"
		} else if water >= 6 && water <= 8 {
			waterStatus = "siaga"
		} else if water > 8 {
			waterStatus = "bahaya"
		}

		if wind <= 6 {
			windStatus = "aman"
		} else if wind >= 7 && wind <= 15 {
			windStatus = "siaga"
		} else if wind > 15 {
			windStatus = "bahaya"
		}

		weather := Weather{
			Water:       water,
			WaterStatus: waterStatus,
			Wind:        wind,
			WindStatus:  windStatus,
		}

		weatherJSON, err := json.Marshal(weather)
		if err != nil {
			fmt.Println("Error marshaling weather object to JSON:", err)
			continue
		}

		fmt.Println(string(weatherJSON))
	}
}
