package service

import (
	"encoding/json"
	"net/http"
)

// Service service for db operation
type Service struct{}

var fullname = map[string]string{
	"btc": "bitcoin",
	"eth": "ethereum",
}

// using coingecko api
// rateapi: https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=btc"
const apiurl = `https://api.coingecko.com/api/v3/simple/price?`

func getrate(from, to string) float64 {
	path := apiurl + "ids=" + fullname[from] + "&vs_currencies=" + to
	response, err := http.Get(path)
	if err != nil {
		return 0
	}
	var target interface{}

	json.NewDecoder(response.Body).Decode(&target)

	// response format
	// {"ethereum":{"btc":0.0312016}}
	// get the value of vs currency
	return target.(map[string]interface{})[fullname[from]].(map[string]interface{})[to].(float64)
}
