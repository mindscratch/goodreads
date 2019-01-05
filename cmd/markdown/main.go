package main

import (
	"os"
	"sort"
	"text/template"

	"github.com/mindscratch/goodreads"
)

func main() {
	records, err := goodreads.ReadFile("data.csv")
	if err != nil {
		panic(err)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].DateRead.Unix() > records[j].DateRead.Unix()
	})

	// Name    | Age
	// --------|------
	// Bob     | 27
	// Alice   | 23
	tmpl, err := template.New("test").Parse(`
Title | Author | Date Read
------|--------|-----------
{{ range . -}}
[{{ .Title }}](https://goodreads.com/book/isbn/{{ .ISBN13 }}) | {{ .Author }} | {{ .DateRead }}
{{ end }}
`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, records)
	if err != nil {
		panic(err)
	}

	// for _, record := range records {
	// 	fmt.Println(record.Title, record.ISBN13)
	// }
}
