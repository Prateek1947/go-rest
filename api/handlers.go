package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	// var movie Movie
	// tx, _ := db.Begin()
	// defer tx.Rollback()
	// stmt, err := tx.Prepare("select * from Movies")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// rows, err := stmt.Query()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var x []string
	// var y []Rating
	// for rows.Next() {
	// 	rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Rated, &movie.Release, &movie.Actors, &movie.Ratings)
	// 	movies = append(movies, movie)
	// }
	// tx.Commit()

	var movies []Movie
	db.Find(&movies)
	for i, movie := range movies {
		db.Model(movie).Related(&movies[i].Actors, "ActorID").Related(&movies[i].Ratings, "RatingID")
	}
	// print(movies[0].Actors[0].Name)
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(movie)
	data, _ := json.Marshal(movies)
	w.Write(data)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println(id)
	var mov Movie
	db.First(&mov, id)
	// if movie != nil {
	json.NewEncoder(w).Encode(movie)
	// }
	// Where(User{Name: "jinzhu"})
}
