package server

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Locations struct {
	Index []Location `json:"index"`
}

type Unique struct {
	UniqueLoc  []string
	CountryLoc map[string][]string
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type AllData struct {
	Artists   []Artist
	Locations Locations
	Unique
}

type AllArtistInfo struct {
	Artist
	Relation
	Coordinates
}

// search

type SearchData struct {
	SearchTerm     string
	SearchedArtist []Artist
	Error          bool
}

type FiltersData struct {
	AllData
	SearchData
}

// map

type Coordinates struct {
	Coo     []string
	MapLink string
}

type Feature struct {
	Center []float64 `json:"center"`
}

type GeocodeResponse struct {
	Features []Feature `json:"features"`
}
