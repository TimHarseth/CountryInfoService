package models

// Struct to represent country information from the first API
type Country_info struct {
	Name struct {
		Common string `json:"common"` // Common name of the country
	} `json:"name"`
	Continent  []string          `json:"continents"` // List of continents the country belongs to
	Population int               `json:"population"` // Population of the country
	Languages  map[string]string `json:"languages"`  // Languages spoken in the country
	Borders    []string          `json:"borders"`    // List of bordering countries
	Capital    []string          `json:"capital"`    // List of capital cities
	Flags      struct {
		PNG string `json:"png"` // URL to the country's flag in PNG format
	} `json:"flags"`
}

// Struct to represent city information from the second API
type City_info struct {
	Cities []string `json:"data"` // List of cities in the country
}

// Struct to combine country and city information for the final response
type Combined_info struct {
	Name       string            `json:"name"`       // Common name of the country
	Continent  []string          `json:"continents"` // List of continents the country belongs to
	Population int               `json:"population"` // Population of the country
	Languages  map[string]string `json:"languages"`  // Languages spoken in the country
	Borders    []string          `json:"borders"`    // List of bordering countries
	Flag       string            `json:"flag"`       // URL to the country's flag in PNG format
	Capital    string            `json:"capital"`    // Capital city of the country
	Cities     []string          `json:"cities"`     // List of cities in the country
}

/*
	type Population_info struct {
		Data struct {
			PopulationValues []map[string]int `json:"populationCounts"`
		} `json:"data"`
	}
*/
/*type Population_info struct {
	Data struct {
		PopulationValues []struct {
			Year  int `json:"year"`
			Value int `json:"value"`
		} `json:"populationCounts"`
	} `json:"data"`
}

type FinalPopulation struct {
	Mean   int             `json:"mean"`
	Values  Population_info.Data.PopulationValues `json:"values"`
}*/
// Definer PopulationValue separat
type PopulationValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type Population_info struct {
	Data struct {
		PopulationValues []PopulationValue `json:"populationCounts"`
	} `json:"data"`
}

type FinalPopulation struct {
	Mean   int               `json:"mean"`
	Values []PopulationValue `json:"values"`
}
