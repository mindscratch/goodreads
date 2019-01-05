package goodreads

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

type DateTime struct {
	time.Time
}

// Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006/01/02"), nil
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	if strings.TrimSpace(csv) == "" {
		return nil
	}
	date.Time, err = time.Parse("2006/01/02", csv)
	return err
}

func (date *DateTime) String() string {
	if date != nil {
		return date.Time.Format("2006/01/02")
	}
	return ""
}

type StringList []string

// Convert the internal []string as CSV string
func (l StringList) MarshalCSV() (string, error) {
	return strings.Join(l, ","), nil
}

// Convert the CSV string as internal []string
func (l StringList) UnmarshalCSV(csv string) (err error) {
	l = strings.Split(csv, ",")
	return nil
}

type OptionalInt struct {
	val int
	set bool
}

// Convert the internal value as CSV string
func (l *OptionalInt) MarshalCSV() (string, error) {
	if l == nil || !l.set {
		return "", nil
	}
	return strconv.Itoa(l.val), nil
}

// Convert the CSV string as internal int
func (l *OptionalInt) UnmarshalCSV(csv string) (err error) {
	csv = strings.Replace(csv, "=", "", -1)
	csv = strings.Replace(csv, "\"", "", -1)
	if csv == "" {
		return nil
	}

	val, err := strconv.Atoi(csv)
	if err != nil {
		return err
	}
	l.val = val
	l.set = true
	return nil
}

func (l *OptionalInt) String() string {
	if l.set {
		return strconv.Itoa(l.val)
	}
	return ""
}

type Record struct {
	BookId                   int         `csv:"Book Id"`
	Title                    string      `csv:"Title"`
	Author                   string      `csv:"Author"`
	AuthorLF                 string      `csv:"Author l-f"`
	AdditionalAuthors        StringList  `csv:"Additional Authors"`
	ISBN                     string      `csv:"ISBN"`
	ISBN13                   OptionalInt `csv:"ISBN13"`
	MyRating                 float32     `csv:"My Rating"`
	AverageRating            float32     `csv:"Average Rating"`
	Publisher                string      `csv:"Publisher"`
	Binding                  string      `csv:"Binding"`
	NumberOfPages            int         `csv:"Number of Pages"`
	YearPublished            int         `csv:"Year Published"`
	OriginalPublicationYear  int         `csv:"Original Publication Year"`
	DateRead                 DateTime    `csv:"Date Read"`
	DateAdded                DateTime    `csv:"Date Added"`
	Bookshelves              StringList  `csv:"Bookshelves"`
	BookshelvesWithPositions StringList  `csv:"Bookshelves with positions"`
	ExclusiveShelf           string      `csv:"Exclusive Shelf"`
	MyReview                 string      `csv:"My Review"`
	Spoiler                  string      `csv:"Spoiler"`
	PrivateNotes             string      `csv:"Private Notes"`
	ReadCount                int         `csv:"Read Count"`
	RecommendedFor           string      `csv:"Recommended For"`
	RecommendedBy            string      `csv:"Recommended By"`
	OwnedCopies              int         `csv:"Owned Copies"`
	OriginalPurchaseDate     DateTime    `csv:"Original Purchase Date"`
	OriginalPurchaseLocation DateTime    `csv:"Original Purchase Location"`
	Condition                string      `csv:"Condition"`
	ConditionDescription     string      `csv:"Condition Description"`
	BCID                     string      `csv:"BCID"`
}

// ReadFile reads the file and converts each row into a Record, an error is returned
// if a problem occurs.
func ReadFile(filename string) ([]*Record, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	records := []*Record{}

	if err := gocsv.UnmarshalFile(f, &records); err != nil {
		return nil, err
	}

	return records, nil
}
