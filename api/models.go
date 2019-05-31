package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Rating movie ratings by different sources
type Rating struct {
	gorm.Model `json:"-"`
	RatingID   uint   `json:"-"`
	Source     string `json:"source"`
	Value      string `json:"value"`
}

//Movie data of various movies
type Movie struct {
	gorm.Model `json:"-"`
	Title      string    `json:"title"`
	Year       string    `json:"year"`
	Rated      string    `json:"rated"`
	Release    time.Time `json:"release"`
	Actors     []Actor   `json:"actors" gorm:"foreignkey:ActorID"`
	Ratings    []Rating  `json:"ratings" gorm:"foreignkey:RatingID"`
}

//Actor data of actors
type Actor struct {
	ActorID uint   `json:"-"`
	Name    string `json:"name"`
}
