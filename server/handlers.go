package server

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var data AllData

// PathHandler : handle every path in a switch
func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		MainPage(w, r)

	case "/artists/":
		ArtistInfo(w, r)

	case "/all_artists_search":
		SearchPage(w, r)

	case "/all_artists":
		AllArtists(w, r)

	default:
		ErrorPage(w, "404 Page not found", http.StatusNotFound)
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	data.getAllData(w, r)
	RenderTemplate(w, "home", data)
}

func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	artists, err1 := GetAllArtists()

	if id <= 0 || id > len(artists) || err != nil {
		err := "404 Page not found"
		ErrorPage(w, err, http.StatusNotFound)
		return
	}

	artistInfo, err2 := GetOneArtistInfo(id)
	relation, err3 := GetRelation(id)

	var coordinates Coordinates

	for key := range relation.DatesLocations {
		loc := strings.Replace(strings.Split(key, " - ")[0], " ", "_", -1)
		coord, error := GetCoords(loc)
		Errors(w, r, error, nil, nil)
		coordinates.Coo = append(coordinates.Coo, coord)
	}

	coordinates.MapLink = SetMarkers(coordinates.Coo)

	Errors(w, r, err1, err2, err3)

	data := AllArtistInfo{artistInfo, relation, coordinates}

	//fmt.Println(coordinates.MapLink)

	RenderTemplate(w, "artist", data)
}

func AllArtists(w http.ResponseWriter, r *http.Request) {
	if len(data.Artists) < 52 {
		data.getAllData(w, r)
	}
	var searchData SearchData
	searchData.Error = false
	allArtistData := FiltersData{data, searchData}
	RenderTemplate(w, "allArtists", allArtistData)
}

func SearchPage(w http.ResponseWriter, r *http.Request) {
	if len(data.Artists) < 52 {
		data.getAllData(w, r)
	}
	var searchData SearchData
	emptyFilters := map[string][]string{"dateRange": {"1900", "2023", "1900", "2023"}, "location": {"-- Select a location --"}, "members": {"", "", "", "", "", "", "", ""}}
	data.getCountryLoc()
	search := r.FormValue("Search")

	if len(search) != 0 && mapsEqual(GetAdvancedSearchInput(r), emptyFilters) {
		var err error
		ids := Search(data, search)
		searchData, err = getSearchResults(ids)
		searchData.SearchTerm = search
		Errors(w, r, err, nil, nil)
	}

	if !mapsEqual(GetAdvancedSearchInput(r), emptyFilters) && len(search) == 0 {
		var err error
		searchData, err = GetFilteredSearch(r, data)
		Errors(w, r, err, nil, nil)
	}

	if len(search) != 0 && !mapsEqual(GetAdvancedSearchInput(r), emptyFilters) {
		var err error
		searchData, err = FullSearch(r, data)
		Errors(w, r, err, nil, nil)
	}

	searchData.Error = len(searchData.SearchedArtist) == 0

	allArtistData := FiltersData{data, searchData}
	RenderTemplate(w, "allArtists", allArtistData)
}

func ErrorPage(w http.ResponseWriter, errors string, code int) {
	w.WriteHeader(code)
	RenderTemplate(w, "error", errors)
}

// RenderTemplate : helps to render the html templates
func RenderTemplate(w http.ResponseWriter, tmplName string, data any) {
	templateCache, err := CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	tmpl, ok := templateCache[tmplName+".gohtml"]

	if !ok {
		http.Error(w, "The template do not exist", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, data)
	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./frontend/templates/*.gohtml")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layout, err := filepath.Glob("./frontend/templates/*.layout.gohtml")

		if err != nil {
			return cache, err
		}
		if len(layout) > 0 {
			tmpl.ParseGlob("./frontend/templates/*.layout.gohtml")
		}
		cache[name] = tmpl
	}
	return cache, nil
}

func Errors(w http.ResponseWriter, r *http.Request, err1 error, err2 error, err3 error) {
	if err1 != nil || err2 != nil || err3 != nil {
		ErrorPage(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
