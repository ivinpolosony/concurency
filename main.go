package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func worker(id int ,url chan string,result chan string){
	for site := range url{
		fmt.Println(site , "started..")
		time.Sleep(time.Second)
		bodyString := httpGet(site)
		fmt.Println(site, "finished.")
		result <- bodyString

	}
}

func httpGet(site string) string {
	var bodyString string
	fmt.Println("getting url..", site)
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}
	defer resp.Body.Close()
	return bodyString
}

func main() {
	sites := make(chan string, 3)
	results := make(chan string, 3)

	sites <- "http://www.google.com"
	sites <- "http://www.golang.org"
	sites <- "http://www.github.com"

	for w := 1; w <= 3; w++ {
		go worker(w, sites, results)
	}

	defer close(sites)

	for a := 1; a <= 3; a++ {
		fmt.Println(<-results)
	}
}

