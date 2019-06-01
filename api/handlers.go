package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

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
func parseForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Helll")
	r.ParseForm()
	actors := r.Form["actor"]
	sources := r.Form["source"]
	values := r.Form["value"]

	// movie=Movie{Actors:Actors{name:actors}}
	var actorsArr []Actor
	var ratings []Rating
	for _, act := range actors {
		actorsArr = append(actorsArr, Actor{Name: act})
	}
	for i, value := range values {
		ratings = append(ratings, Rating{
			Source: sources[i],
			Value:  value,
		})
	}
	_, err := time.Parse("2006-01-02T15:04:05.000Z", r.FormValue("release")+"T11:45:26.371Z")
	if err != nil {
		log.Fatal(err)
	}
	movie = Movie{
		Actors:  actorsArr,
		Rated:   r.FormValue("rated"),
		Ratings: ratings,
		Release: time.Now(),
		Title:   r.FormValue("title"),
		Year:    r.FormValue("year"),
	}
	db.Create(&movie)
}
func uploadForm(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("../templates/form.html"))
	temp.Execute(w, nil)
}
