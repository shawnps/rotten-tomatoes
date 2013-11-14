package rt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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

type MovieSearchResponse struct {
	Total        int
	Movies       []Movie
	Links        map[string]string
	LinkTemplate string
}

func (r *RottenTomatoes) getRequest(params map[string]string, endpoint string) ([]byte, error) {
	v := url.Values{}
	v.Set("apikey", r.Key)
	for key, val := range params {
		v.Set(key, val)
	}
	searchURL := apiURL + endpoint + "?" + v.Encode()
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (r *RottenTomatoes) SearchMovies(q string) ([]Movie, error) {
	p := map[string]string{"q": q}
	e := "/movies.json"
	resp, err := r.getRequest(p, e)
	if err != nil {
		return nil, err
	}
	var m MovieSearchResponse
	err = json.Unmarshal(resp, &m)
	if err != nil {
		return nil, err
	}
	return m.Movies, nil
}
