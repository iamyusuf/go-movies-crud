package main

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func init() {
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "1234",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "1235",
		Title: "Movie 2",
		Director: &Director{
			FirstName: "Jane",
			LastName:  "Doe",
		},
	})
}

func main() {

}
