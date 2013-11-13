rotten-tomatoes
===============

Rotten Tomatoes API client

## Usage:
```Go
apiKey := "your key here"
rt := rt.RottenTomatoes{apiKey}
movies, err := rt.MovieSearch("Good Will Hunting")
if err != nil {
    println("ERROR: ", err.Error())
}   
for _, m := range movies {
    println(m.Title)
}   
```
