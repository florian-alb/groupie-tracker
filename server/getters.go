package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getAllPossiblesLocations(locations Locations) []string {
	var allPossiblesLocations []string

	for _, locations := range locations.Index {
		for _, loc := range locations.Locations {
			if !contains(allPossiblesLocations, loc) {
				allPossiblesLocations = append(allPossiblesLocations, loc)
			}
		}
	}

	return allPossiblesLocations
}

func (locations *Locations) FormatLocation() {

	var formattedLoc []string
	var formatedIndex Locations

	for _, index := range locations.Index {
		for _, loc := range index.Locations {
			loc = strings.Replace(loc, "_", " ", -1)
			loc = strings.Replace(loc, "-", " - ", -1)

			formattedLoc = append(formattedLoc, loc)
		}
		formatedIndex.Index = append(formatedIndex.Index, Location{index.ID, formattedLoc})
		formattedLoc = nil
	}
	*locations = formatedIndex
}

func (rel *Relation) FormatLocation() {
	var sliceRel [][]string

	for key, value := range rel.DatesLocations {
		sliceRel = append(sliceRel, []string{key}, value)
	}

	for i := 0; i < len(sliceRel); i += 2 {
		sliceRel[i][0] = strings.Replace(sliceRel[i][0], "_", " ", -1)
		sliceRel[i][0] = strings.Replace(sliceRel[i][0], "-", " - ", -1)
	}

	rel.DatesLocations = map[string][]string{}

	for i := 0; i < len(sliceRel); i += 2 {
		rel.DatesLocations[sliceRel[i][0]] = sliceRel[i+1]
	}
}

func GetAllArtists() ([]Artist, error) {
	var artists []Artist
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return artists, err
		//log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return artists, err
		//log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &artists)
	if err != nil {
		return artists, err
		//log.Fatal(err)
	}
	defer response.Body.Close()
	return artists, nil
}

func GetAllLocations() (Locations, error) {
	var locations Locations
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return locations, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return locations, err
	}
	err = json.Unmarshal(bytes, &locations)
	if err != nil {
		return locations, err
	}
	defer response.Body.Close()

	locations.FormatLocation()

	return locations, nil
}

func GetRelation(id int) (Relation, error) {
	var relation Relation
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id))
	if err != nil {
		return relation, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return relation, err
	}
	err = json.Unmarshal(bytes, &relation)
	if err != nil {
		return relation, err
	}
	defer response.Body.Close()
	relation.FormatLocation()
	return relation, nil
}

func GetOneArtistInfo(id int) (Artist, error) {
	var artist Artist
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id))
	if err != nil {
		return artist, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return artist, err
	}
	err = json.Unmarshal(bytes, &artist)
	if err != nil {
		return artist, err
	}
	defer response.Body.Close()
	return artist, nil
}

func (ptrData *AllData) getAllData(w http.ResponseWriter, r *http.Request) {
	artists, err1 := GetAllArtists()
	locations, err2 := GetAllLocations()
	Errors(w, r, err1, err2, nil)
	ptrData.Artists, ptrData.Locations, ptrData.UniqueLoc = artists, locations, getAllPossiblesLocations(locations)
}

func (ptrData *AllData) getCountryLoc() {
	countryLoc := make(map[string][]string)
	for _, location := range ptrData.UniqueLoc {
		country := strings.Split(location, " - ")[1]
		var cities []string
		for _, loc := range ptrData.UniqueLoc {
			if strings.Split(loc, " - ")[1] == country {
				cities = append(cities, strings.Split(loc, " - ")[0])
			}
		}

		if _, isKeyExists := countryLoc[country]; isKeyExists {
			continue
		} else {
			countryLoc[country] = cities
		}
	}
	ptrData.CountryLoc = countryLoc
}
