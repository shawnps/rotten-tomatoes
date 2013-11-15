# rt

Rotten Tomatoes API client

Development status: incomplete, many endpoints still missing

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
