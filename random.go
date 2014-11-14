package goplanout

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

func hash(in string) uint64 {

	// Compute 20- byte sha1
	var x [20]byte = sha1.Sum([]byte(in))

	// Get the first 15 characters of the hexdigest.
	var y string = fmt.Sprintf("%x", x[0:8])
	y = y[0 : len(y)-1]

	// Convert hex string into uint64
	var z uint64 = 0
	z, _ = strconv.ParseUint(y, 16, 64)

	return z
}

func generateExperimentId(units interface{}, params map[string]interface{}) string {
	unitstr := generateUnitStr(units)
	salt, exists := params["full_salt"]
	if !exists {
		salt = params["salt"]
	}
	experimentid := salt.(string)
	if unitstr != "" {
		experimentid = experimentid + "." + unitstr
	}
	return experimentid
}

func getHash(m, params map[string]interface{}) uint64 {
	units := evaluate(m["unit"], params)

	_, exists := m["salt"]
	if exists {
		parameter_salt := evaluate(m["salt"], params)
		params["salt"] = parameter_salt.(string)
	}

	experimentid := generateExperimentId(units, params)
	return hash(experimentid)
}

func getUniform(m, params map[string]interface{}, min, max float64) float64 {
	scale, _ := strconv.ParseUint("FFFFFFFFFFFFFFF", 16, 64)
	h := getHash(m, params)
	shift := float64(h) / float64(scale)
	return min + shift*(max-min)
}

type uniformChoice struct{ params map[string]interface{} }

func (s *uniformChoice) execute(m map[string]interface{}) interface{} {
	choices := evaluate(m["choices"], s.params).([]interface{})
	nchoices := uint64(len(choices))
	idx := getHash(m, s.params) % nchoices
	choice := choices[idx]
	return choice
}

type bernoulliTrial struct{ params map[string]interface{} }

func (s *bernoulliTrial) execute(m map[string]interface{}) interface{} {
	pvalue := evaluate(m["p"], s.params).(float64)
	rand_val := getUniform(m, s.params, 0.0, 1.0)
	if rand_val <= pvalue {
		return 0
	}
	return 1
}

type weightedChoice struct{ params map[string]interface{} }

func (s *weightedChoice) execute(m map[string]interface{}) interface{} {
	weights := evaluate(m["weights"], s.params).([]interface{})
	sum, cweights := getCummulativeWeights(weights)
	stop_val := getUniform(m, s.params, 0.0, sum)
	choices := evaluate(m["choices"], s.params).([]interface{})
	for i := range cweights {
		if stop_val <= cweights[i] {
			return choices[i]
		}
	}
	return 0.0
}

type randomFloat struct{ params map[string]interface{} }

func (s *randomFloat) execute(m map[string]interface{}) interface{} {
	min_val := getOrElse(m, "min", 0.0)
	max_val := getOrElse(m, "max", 1.0)
	return getUniform(m, s.params, min_val.(float64), max_val.(float64))
}

type randomInteger struct{ params map[string]interface{} }

func (s *randomInteger) execute(m map[string]interface{}) interface{} {
	min_val := uint64(getOrElse(m, "min", 0.0).(float64))
	max_val := uint64(getOrElse(m, "max", 1.0).(float64))
	return min_val + getHash(m, s.params)%(max_val-min_val+1)
}
