package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "strconv"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "goapi-demo/config"
)

var cfg config.Config
var db *sql.DB
var jokeDb *JokeDb

func showJoke(c web.C, w http.ResponseWriter, r *http.Request) {
    jokeId, err := strconv.Atoi(c.URLParams["jokeId"])
    if err != nil {
        log.Fatal(err)
    }

    joke, err := jokeDb.FetchJoke(jokeId)
    switch {
        case err == sql.ErrNoRows:
            log.Printf("No joke with that ID.")
            fmt.Fprint(w, "404")
        case err != nil:
            log.Fatal(err)
        default:
            sjson, _ := json.Marshal(joke)
            fmt.Fprintf(w, string(sjson))
    }
}

func init() {
    var err error
    
    cfg, err = config.NewFromFile("config.json")
    if err != nil {
        log.Fatal(err)
    }

    db, err = sql.Open("mysql", cfg.MySqlDsn)
    if err != nil {
        log.Fatal(err)
    }

    jokeDb = &JokeDb{db}
}

func main() {
    goji.Use(ContentTypeJson)
    goji.Get("/jokes/:jokeId", showJoke)
    goji.Serve()

    db.Close()
}
