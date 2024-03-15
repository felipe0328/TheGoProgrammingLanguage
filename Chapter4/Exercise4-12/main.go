// Exercise 4.12 The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkdc.com/571/info.0.json produces a detailed description of comic 571, one of many favorites. Download each
// URL (once!) and build and offline index. Write a tool xkcd that, using this index, prints the URL and transcript of
// each comic that matches a search term provided on the command line.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Comic struct {
	Num   int
	URL   string
	Alt   string
	Title string
}

type ComicGetter struct {
	Comics *map[int]Comic
}

const fileName = "comicsIndex.json"

var comicFlag = flag.Int("c", 0, "comic to obtain")

func main() {

	var jsonComics map[int]Comic

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		jsonComics = createJsonIndex()
	} else {
		if err = json.NewDecoder(f).Decode(&jsonComics); err != nil {
			f.Close()
			fmt.Println(err)
			return
		}

		f.Close()
	}

	flag.Parse()

	if *comicFlag == 0 {
		fmt.Println("You should provide a comic index to search")
		return
	}

	comic := jsonComics[*comicFlag]
	fmt.Printf("Title: %s\nUrl: %s\nNumber: %d\nDescription:%s\n", comic.Title, comic.URL, comic.Num, comic.Alt)

}

func createJsonIndex() map[int]Comic {

	ch := make(chan Comic)

	jsonResult := make(map[int]Comic, 0)
	for i := 1; i <= 1000; i++ {
		i := i
		go getComic(i, ch)
	}

	for i := 1; i <= 1000; i++ {
		fmt.Println("Appending comic ", i)
		comic := <-ch
		jsonResult[comic.Num] = comic
	}

	result, err := json.Marshal(jsonResult)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Creating file")
	os.WriteFile(fileName, result, os.FileMode(0777))
	fmt.Println("File Created")

	return jsonResult
}

func getComic(index int, ch chan Comic) {
	if index == 404 { //404 does not exists
		ch <- Comic{
			URL: "404",
		}
		return
	}

	fmt.Println("Getting comic: ", index)
	url := "https://xkcd.com/%d"
	info := "/info.0.json"
	comicUrl := fmt.Sprintf(url, index)
	resp, err := http.Get(fmt.Sprintf("%s%s", comicUrl, info))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return
	}

	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		fmt.Println(err)
		return
	}

	comic.URL = comicUrl

	ch <- comic
}
