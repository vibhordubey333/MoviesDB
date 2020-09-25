package main

import (
	"MoviesDB/entities"
	"MoviesDB/service"
	"fmt"
)

func main() {
	movieName := "Hannibal"
	searchResult, err := service.ReadDocument(entities.IMDBRegistry{MovieName: movieName}, &entities.IMDBRegistry{})
	fmt.Println("Search Result : ", searchResult, " err: ", err)
}
