// Exercise 4.13 The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/ for a movie
// by name and download its poster image. Write a tool poster that download downloads the poster image for the movie
// named on the command line
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OmdbApiResponse struct {
	Title  string
	Poster string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to provide a movie title to search")
		return
	}

	movieTitle := os.Args[1]
	poster(movieTitle)
}

func poster(title string) {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://omdbapi.com/?t=%s&apikey=%s", title, apiKey)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Unable to get movie, ", err.Error())
		return
	}

	defer resp.Body.Close()

	var info OmdbApiResponse
	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		fmt.Println("Unable to decode info", err)
		return
	}

	imageName := fmt.Sprintf("%s.jpg", info.Title)
	f, err := os.Create(imageName)
	if err != nil {
		fmt.Println("Unable to create image: ", imageName, err)
		return
	}

	defer f.Close()
	img, err := http.Get(info.Poster)
	if err != nil {
		fmt.Println("Unable to get image", err)
		return
	}

	defer img.Body.Close()
	_, err = io.Copy(f, img.Body)
	if err != nil {
		fmt.Println("Unable to save image ", err)
	}
}
