// Exercise 7.9 Use the html/template package to replace printTracks with a function that
// displays the tracks as an HTML table. Use the solution to the previous exercise to arrange
// that each click on a column head makes an HTTP request to sort the table
package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"text/template"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Marting Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func formatTracks() string {
	return `
	<html>
		<head>
		<head>
	</html>`
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	fmt.Fprint(tw, "\n\n")
	tw.Flush()
}

type customSort struct {
	t                 []*Track
	lessFunctionsHead *linkedList
	lessFunctionsTail *linkedList
}

func (x customSort) Len() int      { return len(x.t) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x customSort) Less(i, j int) bool {
	for lt := x.lessFunctionsHead.next; lt != nil; lt = lt.next {
		if lt.validator(i, j) {
			return lt.less(i, j)
		}
	}
	return false
}

type linkedList struct {
	validator func(i, j int) bool
	less      func(i, j int) bool
	prev      *linkedList
	next      *linkedList
}

func (x customSort) canSortByTitle(i, j int) bool { return x.t[i].Title != x.t[j].Title }
func (x customSort) sortByTitle(i, j int) bool    { return x.t[i].Title < x.t[j].Title }
func (x customSort) canSortByYear(i, j int) bool  { return x.t[i].Year != x.t[j].Year }
func (x customSort) sortByYear(i, j int) bool     { return x.t[i].Year < x.t[j].Year }

func (x customSort) sendToHead(i *linkedList) {
	if x.lessFunctionsHead.next == i {
		return
	}

	if i.next != nil {
		i.next.prev = i.prev
	}

	if i.prev != nil {
		i.prev.next = i.next
	}

	i.prev = x.lessFunctionsHead
	i.next = x.lessFunctionsHead.next

	x.lessFunctionsHead.next.prev = i
	x.lessFunctionsHead.next = i
}

var filters = make(map[string]*linkedList)

func main() {
	cs := customSort{tracks, &linkedList{}, &linkedList{}}

	filters["Title"] = &linkedList{validator: cs.canSortByTitle, less: cs.sortByTitle}
	filters["Year"] = &linkedList{validator: cs.canSortByYear, less: cs.sortByYear}

	cs.lessFunctionsHead.next = filters["Title"]
	filters["Title"].next = filters["Year"]
	filters["Year"].next = cs.lessFunctionsTail

	http.HandleFunc("/", cs.ShowTableInfo)
	panic(http.ListenAndServe(":8080", nil))
}

func (x customSort) ShowTableInfo(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")

	if filter != "" {
		if filterObject, ok := filters[filter]; ok {
			x.sendToHead(filterObject)

			if sort.IsSorted(x) {
				sort.Sort(sort.Reverse(x))
			} else {
				sort.Sort(x)
			}

		}
	}

	report := template.Must(template.New("trackInfo").Parse(templateToUse))

	if err := report.Execute(w, tracks); err != nil {
		fmt.Println("Error creating template: ", err)
	}
}
