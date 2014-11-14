package goplanout

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"fmt"
)

func TestUniformChoice(t *testing.T) {

	// Read planout code in json from file on disk.
	// x = uniformChoice(choices=[1,2,3,4], unit=[userid, country]);
	f := "test/unit.txt"
	data, _ := ioutil.ReadFile(f)

	// Parse json into map[string]interface{}
	var js map[string]interface{}
	json.Unmarshal(data, &js)

	// Set input parameters
	id1 := generateString()
	id2 := generateString()
	id3 := id1

	// Set inputs for the experiment.
	params := make(map[string]interface{})
	params["country"] = "US"
	params["full_salt"] = "experiment_salt"

	// Run experiment with id1
	params["userid"] = id1
	Experiment(js, params)
	x1 := params["x"]
	fmt.Printf("Params: %v\n", params)

 	// Run experiment with id2
 	params["userid"] = id2
 	Experiment(js, params)
	fmt.Printf("Params: %v\n", params)

	// Run experiment with id3
	params["userid"] = id3
	Experiment(js, params)
	x3 := params["x"]
	fmt.Printf("Params: %v\n", params)

	if compare(x1, x3) != 0 {
		t.Errorf("Expected same choice for [userid %v, x %v] and [userid %v, x %v]\n", id1, x1, id3, x3)
	}
}
