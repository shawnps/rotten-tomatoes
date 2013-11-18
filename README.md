# rt

Rotten Tomatoes API client

Development status: incomplete, many endpoints still missing

## Usage:
```Go
package main

import (
    "fmt"
    "github.com/shawnps/rt"
)

func main() {
    apiKey := "your key here"
    rt := rt.RottenTomatoes{apiKey}
    movies, err := rt.SearchMovies("Good Will Hunting")
    if err != nil {
        fmt.Println("ERROR: ", err.Error())
    }   
    for _, m := range movies {
        fmt.Println(m)
    }   
}
```
