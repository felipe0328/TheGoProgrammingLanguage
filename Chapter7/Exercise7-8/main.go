// Exercise 7.8 Many GUIs provide a table widget with a stateful multi-tier sort: The primary
// sort key is the most recently clicked column head, the secondary sort key is the second-most
// recently clicked column head, and so on. Define an implementation of sort.Interface for use
// by such a table. Compare that approach with repeated sorting using sort.Stable.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

func main() {
	printTracks(tracks)

	cs := customSort{tracks, &linkedList{}, &linkedList{}}

	filters := make(map[string]*linkedList)
	filters["Title"] = &linkedList{validator: cs.canSortByTitle, less: cs.sortByTitle}
	filters["Year"] = &linkedList{validator: cs.canSortByYear, less: cs.sortByYear}

	cs.lessFunctionsHead.next = filters["Title"]
	filters["Title"].next = filters["Year"]
	filters["Year"].next = cs.lessFunctionsTail

	sort.Sort(cs)
	printTracks(tracks)
	cs.sendToHead(filters["Year"])
	sort.Sort(cs)
	printTracks(tracks)
}
