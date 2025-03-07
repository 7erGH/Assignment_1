package models

type CountryInfo struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

type PopulationData struct {
	Mean   int               `json:"mean"`
	Values []PopulationValue `json:"values"`
}

type PopulationValue struct {
	Year  int    `json:"year"`
	Value int    `json:"value"`
	Sex   string `json:"sex,omitempty"`
}

type StatusUpdate struct {
	CountriesNowAPI  int    `json:"countriesnowapi"`
	RestCountriesAPI int    `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

type RestCountries struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		PNG string `json:"png"`
	} `json:"flags"`
	Capital []string `json:"capital"`
}

// from countriesNow API
type CountriesNow struct {
	Country string `json:"country"`
}


type CountriesNCities struct {
	Error bool     `json:"error"`
	Msg   string   `json:"msg"`
	Data  []string `json:"data"`
}

// response structure for population data from countriesNow API
type CountriesNPopulation struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Country    string            `json:"country"`
		Code       string            `json:"code"`
		Population []PopulationValue `json:"populationCounts"`
	} `json:"data"`
}

