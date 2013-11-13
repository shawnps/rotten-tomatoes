package rt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiURL = "http://api.rottentomatoes.com/api/public/v1.0"
)

type RottenTomatoes struct {
	Key string
}

type Ratings struct {
	CriticsRating  string `json:"critics_rating"`
	CriticsScore   int    `json:"critics_score"`
	AudienceRating string `json:"audience_rating"`
	AudienceScore  int    `json:"audience_score"`
}
type Actor struct {
	Name       string
	Id         string
	Characters []string
}

type Movie struct {
	Id               string
	Title            string
	Year             int
	MPAARating       string `json:"mpaa_rating"`
	Runtime          int
	CriticsConsensus string            `json:"critics_consensus"`
	ReleaseDates     map[string]string `json:"release_dates"`
	Ratings          Ratings
	Synopsis         string
	Posters          map[string]string
	AbridgedCast     []Actor           `json:"abridged_cast"`
	AlternateIds     map[string]string `json:"alternate_ids"`
	Links            map[string]string
}
