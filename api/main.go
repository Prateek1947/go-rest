package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var movie Movie
var actor Actor
var rating Rating
var db *gorm.DB

func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("../media"))
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", fs))
	r.HandleFunc("/movies/{something}", getAllMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/uploadMovies", uploadForm).Methods(http.MethodGet)
	r.HandleFunc("/uploadMovies", parseForm).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)
}
func init() {
	// movies = append(movies, Movie{1, "Guardians of the Galaxy Vol. 2", "2017", "PG-13", time.Now(), []string{"Chris Pratt", "Dave Bautista", "Vin Diesel"}, []Rating{Rating{"Internet Movie Database", "7.7/10"}}})
	var err error
	db, err = gorm.Open("mysql", "prateek:prateek@1@tcp(127.0.0.1:3306)/PrateekDatabase?parseTime=true")
	// defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&rating, &movie, &actor)
	if !db.HasTable(&rating) || !db.HasTable(&movie) || !db.HasTable(&actor) {
		db.CreateTable(&rating)
		db.CreateTable(&movie)
		db.CreateTable(&actor)
	}

	movie = Movie{Title: "xzxnjdjsflg",
		Year:    "2017",
		Rated:   "PG-13sdasf",
		Release: time.Now(),
		Actors: []Actor{
			Actor{Name: "Prateek"},
			Actor{Name: "Roman Reigns"},
		},
		Ratings: []Rating{
			Rating{Source: "dfdgf Movie Database", Value: "7.7/10"},
			Rating{Source: "IMDBsaddsgd", Value: "1.0/10"},
			Rating{Source: "sadsfdsgfd", Value: "7.7/10"},
		},
	}
	// var movies []Movie
	// db.Create(&movie)
	// db.Find(&movies)
	// for _, movie := range movies {
	// 	db.Model(&movie).Related(&movie.Ratings, "RatingID").Related(&movie.Actors, "ActorID")
	// 	fmt.Println(movie)
	// }
}
