package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "233512",
		Title: "Life of Pi",
		Director: &Director{
			Firstname: "Aang",
			Lastname:  "Lee",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "345678",
		Title: "The Shawshank Redemption",
		Director: &Director{
			Firstname: "Frank",
			Lastname:  "Darabont",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "987654",
		Title: "The Godfather",
		Director: &Director{
			Firstname: "Francis",
			Lastname:  "Ford Coppola",
		},
	})
	movies = append(movies, Movie{
		ID:    "4",
		Isbn:  "456789",
		Title: "The Dark Knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	})
	movies = append(movies, Movie{
		ID:    "5",
		Isbn:  "123456",
		Title: "Pulp Fiction",
		Director: &Director{
			Firstname: "Quentin",
			Lastname:  "Tarantino",
		},
	})
	movies = append(movies, Movie{
		ID:    "6",
		Isbn:  "789012",
		Title: "The Lord of the Rings: The Return of the King",
		Director: &Director{
			Firstname: "Peter",
			Lastname:  "Jackson",
		},
	})
	movies = append(movies, Movie{
		ID:    "7",
		Isbn:  "345679",
		Title: "Schindler's List",
		Director: &Director{
			Firstname: "Steven",
			Lastname:  "Spielberg",
		},
	})
	movies = append(movies, Movie{
		ID:    "8",
		Isbn:  "987655",
		Title: "The Good, the Bad and the Ugly",
		Director: &Director{
			Firstname: "Sergio",
			Lastname:  "Leone",
		},
	})
	movies = append(movies, Movie{
		ID:    "9",
		Isbn:  "456790",
		Title: "The Silence of the Lambs",
		Director: &Director{
			Firstname: "Jonathan",
			Lastname:  "Demme",
		},
	})
	movies = append(movies, Movie{
		ID:    "10",
		Isbn:  "123457",
		Title: "Star Wars: Episode IV - A New Hope",
		Director: &Director{
			Firstname: "George",
			Lastname:  "Lucas",
		},
	})
	movies = append(movies, Movie{
		ID:    "11",
		Isbn:  "789013",
		Title: "The Matrix",
		Director: &Director{
			Firstname: "The Wachowskis",
			Lastname:  "",
		},
	})
	movies = append(movies, Movie{
		ID:    "12",
		Isbn:  "345680",
		Title: "Fight Club",
		Director: &Director{
			Firstname: "David",
			Lastname:  "Fincher",
		},
	})
	movies = append(movies, Movie{
		ID:    "13",
		Isbn:  "987656",
		Title: "Casablanca",
		Director: &Director{
			Firstname: "Michael",
			Lastname:  "Curtiz",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
