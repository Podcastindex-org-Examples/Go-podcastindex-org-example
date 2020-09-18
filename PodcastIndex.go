package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strconv"
	"crypto/sha1"
)

func main() {
	fmt.Println("running...")
	// ======== Required values ======== 
	// WARNING: don't publish these to public repositories or in public places!
	// NOTE: values below are sample values, to get your own values go to https://api.podcastindex.org 
	var apiKey string = "ABC"
	var apiSecret string = "ABC"
	// prep for crypto
	now	:= time.Now()
	var apiHeaderTime string = strconv.FormatInt(now.Unix(), 10)	
	var data4Hash string = apiKey + apiSecret + apiHeaderTime	
	fmt.Println("data4Hash: "+data4Hash)
	// ======== Hash them to get the Authorization token ========
	h := sha1.New()	
	h.Write([]byte(data4Hash))
	hash := h.Sum(nil)
	var hashString string = fmt.Sprintf("%x", hash)
	fmt.Println("hashString: "+hashString)
	// ======== Send the request and collect/show the results ======== 
	var query string = "bastiat"		
	url := "https://api.podcastindex.org/api/1.0/search/byterm?q="+query

	spaceClient := http.Client{
		Timeout: time.Second * 33, 
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "SuperPodcastPlayer/1.8")
	req.Header.Set("X-Auth-Date", apiHeaderTime)
	req.Header.Set("X-Auth-Key", apiKey)
	req.Header.Set("Authorization", hashString)

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println("body: \n"+string(body))
}

