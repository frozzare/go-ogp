package ogp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Fetch all Open Graph meta tags and return a list
func Fetch(url string) map[string]interface{} {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	r, _ := regexp.Compile("<meta.*property=\"og:(.*)\".*content=\"(.*)\".*>")
	matches := r.FindAllStringSubmatch(string(body), -1)

	list := make(map[string]interface{})

	for i := range matches {
		list[matches[i][1]] = matches[i][2]
	}

	return list
}

// Print out all Open Graph meta tags in the terminal
func Print(url string) {
	list := Fetch(url)

	if list["title"] != nil {
		fmt.Printf("\n\x1b[33;1m%s\x1b[0m\n", list["title"])
	}

	for key, value := range list {
		fmt.Printf("\n\x1b[32;1mog:%s\x1b[0m\n%s\n", key, value)
	}

	fmt.Printf("\n")
}
