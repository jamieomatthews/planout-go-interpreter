package main

import (
	"fmt"
	"crypto/sha1"
	"strconv"
)

var salts map[string]string

func getHash(in string) uint64 {

	// full_salt := getFullSalt()
	// salt := getSalt()
	// experiment_salt := getExperimentSalt()

	// if len(fullsalt) == 0
	// Compute 20- byte sha1
	var x [20]byte = sha1.Sum([]byte(in))

	// Get the first 15 characters of the hexdigest.
	var y string = fmt.Sprintf("%x", x[0:8])
	y = y[0:len(y)-1]

	// Convert hex string into uint64
	var z uint64 = 0
	z, _ = strconv.ParseUint(y, 16, 64)

	return z
}

func getUniform(in string, min, max float64) float64 {
	scale,_ := strconv.ParseUint("FFFFFFFFFFFFFFF", 16, 64)
	h := getHash(in)
	shift := float64(h)/float64(scale)
	return min + shift*(max-min)
}

func main() {
	fmt.Println(getUniform("Hello", 0.0, 1.0))
}
