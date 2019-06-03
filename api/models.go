package main

import (
	"time"
)

//Rating movie ratings by different sources
type Rating struct {
	RatingID uint   `json:"-"`
	Source   string `json:"source"`
	Value    string `json:"value"`
}

//Movie data of various movies
type Movie struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	Title     string    `json:"title"`
	Year      string    `json:"year"`
	Rated     string    `json:"rated"`
	PosterURI string    `json:"-"`
	Release   time.Time `json:"release"`
	Actors    []Actor   `json:"-" gorm:"foreignkey:ActorID"`
	Ratings   []Rating  `json:"ratings" gorm:"foreignkey:RatingID"`
}

//Actor data of actors
type Actor struct {
	ActorID uint   `json:"-"`
	Name    string `json:"name"`
}
