package server

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func Search(data AllData, search string) []int {
	if strings.Contains(search, " - ") {
		search = strings.Split(search, " - ")[0]
	}
	search = strings.ToLower(search)

	var id []int
	for _, result := range data.Artists {

		for _, name := range strings.Split(result.Name, " ") {
			if strings.Contains(strings.ToLower(name), search) {
				id = append(id, result.ID)
				break
			}
		}

		if strings.Contains(strings.ToLower(result.FirstAlbum), search) || strings.Contains(strings.ToLower(strconv.Itoa(result.CreationDate)), search) {
			id = append(id, result.ID)
		}

		for _, members := range result.Members {
			for _, name := range strings.Split(members, " ") {
				if strings.ToLower(name) == search {
					id = append(id, result.ID)
					break
				}
			}
		}
	}
	for _, result := range data.Locations.Index {
		for _, location := range result.Locations {
			for _, word := range strings.Split(location, " ") {
				if strings.ToLower(word) == search {
					id = append(id, result.ID)
					break
				}
			}
		}
	}
	return removeDuplicates(id)
}

func getSearchResults(ids []int) (SearchData, error) {
	var output SearchData
	if len(ids) == 0 {
		return output, nil
	}
	for _, id := range ids {
		artistInfo, err1 := GetOneArtistInfo(id)

		if err1 != nil {
			return output, err1
		}

		output.SearchedArtist = append(output.SearchedArtist, artistInfo)
	}

	return output, nil
}

func GetAdvancedSearchInput(r *http.Request) map[string][]string {

	advancedFilters := make(map[string][]string)

	advancedFilters["dateRange"] = []string{
		r.FormValue("creation-date-min"),
		r.FormValue("creation-date-max"),
		r.FormValue("first-album-date-min"),
		r.FormValue("first-album-date-max")}

	advancedFilters["members"] = []string{
		r.FormValue("one-member"),
		r.FormValue("two-members"),
		r.FormValue("three-members"),
		r.FormValue("four-members"),
		r.FormValue("five-members"),
		r.FormValue("six-members"),
		r.FormValue("seven-members"),
		r.FormValue("eight-members")}

	advancedFilters["location"] = []string{r.FormValue("location")}

	return advancedFilters
}

func AdvancedSearch(data AllData, filters map[string][]string) []int {
	var mapId = make(map[int][]int)

	for key, value := range filters {

		switch key {
		case "dateRange":
			creationMin, err1 := strconv.Atoi(value[0])
			creationMax, err2 := strconv.Atoi(value[1])
			fistAlbumMin, err1 := strconv.Atoi(value[2])
			fistAlbumMax, err2 := strconv.Atoi(value[3])

			if err1 != nil || err2 != nil {
				log.Fatal(err1, err2)
			}

			for _, artist := range data.Artists {
				firstAlbum, err := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])
				if err != nil {
					log.Fatal(err)
				}

				if artist.CreationDate > creationMin && artist.CreationDate < creationMax && firstAlbum > fistAlbumMin && firstAlbum < fistAlbumMax {
					mapId[artist.ID] = append(mapId[artist.ID], artist.ID)
				}
			}
			break

		case "members":
			for i, checkbox := range value {
				if checkbox == "on" {
					for _, artist := range data.Artists {
						if i+1 == len(artist.Members) {
							mapId[artist.ID] = append(mapId[artist.ID], artist.ID)
						}
					}
				}
			}
			break
		case "location":
			for _, location := range data.Locations.Index {
				for _, loc := range location.Locations {
					if strings.Contains(loc, value[0]) {
						mapId[location.ID] = append(mapId[location.ID], location.ID)
					}
				}
			}
			break
		}
	}

	var id []int
	for key, value := range mapId {
		if filters["location"][0] == "-- Select a location --" {
			if len(value) == 2 {
				id = append(id, key)
			}
		} else {
			if len(value) == 3 {
				id = append(id, key)
			}
		}
	}

	return id
}

func GetFilteredSearch(r *http.Request, data AllData) (SearchData, error) {
	filters := GetAdvancedSearchInput(r)
	id := AdvancedSearch(data, filters)
	sort.Ints(id)

	searchData, error := getSearchResults(id)
	searchData.Error = len(id) >= 1

	return searchData, error
}

func FullSearch(r *http.Request, data AllData) (SearchData, error) {
	filters := GetAdvancedSearchInput(r)
	id := AdvancedSearch(data, filters)
	id = append(id, Search(data, r.FormValue("Search"))...)
	id = findDuplicates(id)

	searchData, error := getSearchResults(id)
	searchData.Error = len(id) <= 1
	return searchData, error
}
