development status: incomplete, many endpoints still missing

rt
===============

Rotten Tomatoes API client

## Usage:
```Go
import "github.com/shawnps/rt"

apiKey := "your key here"
rt := rt.RottenTomatoes{apiKey}
movies, err := rt.SearchMovies("Good Will Hunting")
if err != nil {
    println("ERROR: ", err.Error())
}   
for _, m := range movies {
    println(m.Title)
}   
```
