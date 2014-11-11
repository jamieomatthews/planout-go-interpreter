package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Read planout code in json from file on disk.
	f := "test/" + os.Args[1]
	data, _ := ioutil.ReadFile(f)

	// Parse json into map[string]interface{}
	var js map[string]interface{}
	json.Unmarshal(data, &js)

	// Set input parameters
	params := make(map[string]interface{})
	params["country"] = "US"

	// Run experiment
	success := experiment(js, params)
	if success {
		fmt.Println("Success running experiment!")
	} else {
		fmt.Println("Failed running experiment!")
	}
}
