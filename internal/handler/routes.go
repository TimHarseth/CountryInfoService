package handler

import (
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `
			<h1>Welcome!</h1>
			<p>Available endpoints:</p>
			<ul>
				<li><a href="/countryinfo/v1/info/">/info</a> - Shows info about specific countries</li>
				<li><a href="/countryinfo/v1/population/">/population</a> - Shows population data across several years</li>
				<li><a href="/countryinfo/v1/status/">/status</a> - Shows status and uptime of application</li> 
			</ul>
		`)
}

func DefaultInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `
			<h1>Welcome!</h1>
			<p>Guide:</p>
			<p>GET info/{:two_letter_country_code}{?limit=10}</p>
			<p> Example: <a href=/countryinfo/v1/info/no?limit=5>link</a> - /countryinfo/v1/info/no?limit=5
		`)
}
func DefaultPopulation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `
			<h1>Welcome!</h1>
			<p>Guide:</p>
			<p>GET population/{:two_letter_country_code}{?limit={:startYear-endYear}}</p>
			<p> Example: <a href=/countryinfo/v1/population/no?limit=2010-2015>link</a> - /countryinfo/v1/population/no?limit=2010-2015
		`)
}
