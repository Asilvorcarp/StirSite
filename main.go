package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

const (
	// website to stir (default: localhost:3000)
	WEBSITE = "http://localhost:3000"
	// path to sitemap (default: sitemap.xml)
	SITEMAP = "sitemap.xml"
	// interval between each stir (default: 1 hour)
	INTERVAL = time.Hour
	// number of consecutive requests to make each time (default: 3)
	CONSECUTIVE = 3
	// number of seconds to wait before making a request (default: 1 second)
	DELAY = time.Second
)

type Sitemap struct {
	Locations []string `xml:"url>loc"`
}

// get all urls from sitemap
func getAllUrls() []string {
	// get sitemap as xml
	resp, err := http.Get(WEBSITE + "/" + SITEMAP)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// parse xml
	var sitemap Sitemap
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&sitemap)
	if err != nil {
		fmt.Println(err)
	}
	// return urls
	return sitemap.Locations
}

// make a number of requests to a url
func StirURL(url string) {
	for i := 0; i < CONSECUTIVE; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()
		time.Sleep(DELAY)
	}
	fmt.Println("- Stirred", url)
}

// sitr the website once
func StirOnce() {
	urls := getAllUrls()
	for _, url := range urls {
		StirURL(url)
	}
	fmt.Println("- Stirred", len(urls), "urls")
}

func main() {
	for {
		// print time
		fmt.Println("> Stir at", time.Now())
		// make request
		go StirOnce()
		// wait
		time.Sleep(INTERVAL)
	}
}
