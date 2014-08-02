package main

import (
    "database/sql"
)

type Joke struct {
    Id int          `json:"id"`
    Text string     `json:"text"`
    Likes int       `json:"likes"`
    Favourites int  `json:"favourites"`
}

type JokeDb struct {
    db *sql.DB
}

func (jokeDb JokeDb) FetchJoke(id int) (joke *Joke, err error) {
    joke = new(Joke)
    row := db.QueryRow("SELECT id, text, likes, favourites FROM jokes WHERE id=?", id)
    err = row.Scan(&joke.Id, &joke.Text, &joke.Likes, &joke.Favourites)
    return
}
