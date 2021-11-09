package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

type ByColumns struct {
	t          []*Track
	columns    []columnCmp
	maxColumns int
}

type columnCmp func(x, y *Track) comparison

func (x ByColumns) Len() int           { return len(x.t) }
func (x ByColumns) Less(i, j int) bool {
	for _, f := range x.columns {
		switch f(x.t[i], x.t[j]) {
		case lt:
			return true
		case gt:
			return false
		case eq:
			continue
		}
	}
	return false
}
func (x ByColumns) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

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
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//type customSort struct {
//	keyOrder []string
//	t []*Track
//	less func(keyOrder []string, x, y *Track) bool
//}
func (c *ByColumns) LessTitle(a, b *Track) comparison {
	switch {
	case a.Title == b.Title:
		return eq
	case a.Title < b.Title:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) LessArtist(a, b *Track) comparison {
	switch {
	case a.Artist == b.Artist:
		return eq
	case a.Artist < b.Artist:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) Select(cmp columnCmp) {
	c.columns = append([]columnCmp{cmp}, c.columns...)
}

func NewByColumns(p []*Track, maxColumns int) *ByColumns {
	return &ByColumns{p, nil, maxColumns}
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

var html = template.Must(template.New("people").Parse(`
<html>
<body>
<table>
	<tr>
		<th><a href="?sort=title">Title</a></th>
		<th><a href="?sort=artist">Artist</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
	</td>
{{end}}
</body>
</html>
`))

func main() {
	//p := []string{"hello","world"}
	//sort.Stable(p)
	//p := []string{"Artist", "Year", "Title", "Album", "Length"}
	nb := NewByColumns(tracks, 4)

	//c := column.NewByColumns(people, 2)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "artist":
			nb.Select(nb.LessArtist)
		case "title":
			nb.Select(nb.LessTitle)
		}
		sort.Sort(nb)
		err := html.Execute(w, tracks)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

	//fmt.Println("排序前：")
	//printTracks(nb.t)
	//
	//nb.Select(nb.LessTitle)
	//sort.Sort(nb)
	//
	//fmt.Println("按标题排序后：")
	//printTracks(nb.t)
	//
	//nb.Select(nb.LessArtist)
	//sort.Sort(nb)
	//
	//fmt.Println("按歌手排序后：")
	//printTracks(nb.t)

}