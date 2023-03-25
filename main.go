package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// global vars with default values
var (
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

// try to get all the global vars from the environment / .env file
// if they are not set, use the default values
// >> this is run before the main function
func init() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Not loading .env file")
	}
	if website := os.Getenv("WEBSITE"); website != "" {
		fmt.Println("WEBSITE =", website)
		WEBSITE = website
	}
	if sitemap := os.Getenv("SITEMAP"); sitemap != "" {
		fmt.Println("SITEMAP =", sitemap)
		SITEMAP = sitemap
	}
	if interval := os.Getenv("INTERVAL"); interval != "" {
		fmt.Println("INTERVAL =", interval)
		INTERVAL, _ = time.ParseDuration(interval)
	}
	if consecutive := os.Getenv("CONSECUTIVE"); consecutive != "" {
		fmt.Println("CONSECUTIVE =", consecutive)
		CONSECUTIVE, _ = strconv.Atoi(consecutive)
	}
	if delay := os.Getenv("DELAY"); delay != "" {
		fmt.Println("DELAY =", delay)
		DELAY, _ = time.ParseDuration(delay)
	}
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

// stir the website once
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
