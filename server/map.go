package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const ACCESSTOKEN = "pk.eyJ1IjoibmluamF3aWxsNTQzIiwiYSI6ImNsZDRidWdrNjBvbDczcW9jajU5c3UxdXAifQ._jM-ztlL3-V9_0WjlFB01A"

func GetCoords(location string) (string, error) {
	var geocodeResponse GeocodeResponse
	var coo string

	apiUrl := "https://api.mapbox.com/geocoding/v5/mapbox.places/"

	response, err := http.Get(apiUrl + location + ".json?proximity=ip&limit=1&access_token=" + ACCESSTOKEN)

	err = json.NewDecoder(response.Body).Decode(&geocodeResponse)
	if err != nil {
		return coo, err
	}

	if len(geocodeResponse.Features) == 0 {
		fmt.Println("No location found")
		return coo, nil
	}

	lgn := strconv.FormatFloat(geocodeResponse.Features[0].Center[1], 'f', -1, 64)
	lat := strconv.FormatFloat(geocodeResponse.Features[0].Center[0], 'f', -1, 64)

	coo = lat + "," + lgn

	defer response.Body.Close()

	return coo, nil
}

func SetMarkers(coords []string) string {

	var pings string
	for _, coord := range coords {
		pings += "pin-s+1DB954(" + coord + "),"
	}
	pings = pings[:len(pings)-1]
	pings += "/"

	mapURL := "https://api.mapbox.com/styles/v1/mapbox/dark-v10/static/" + pings + "0,0,0.1/1280x1280?access_token=" + ACCESSTOKEN

	return mapURL
}
