package file

import (
	"fmt"
	"time"
)

type Filename struct {
	date string
	ext  string
}

const LAYOUT = "060102"
const EXT = ".md"

func NewFilenameFromDate(date time.Time) Filename {
	return Filename{date: date.Format(LAYOUT), ext: EXT}
}

func NewFilenameFromString(filename string) (Filename, error) {
	layout := LAYOUT + EXT
	date, err := time.Parse(layout, filename)
	if err != nil {
		return Filename{}, fmt.Errorf("error in NewFilenameFromString: %w", err)
	}

	return NewFilenameFromDate(date), nil
}

func NewFilenameToday() Filename {
	date := time.Now()
	return NewFilenameFromDate(date)
}

func (f Filename) Filename() string {
	return f.date + f.ext
}

func (f Filename) Date() (time.Time, error) {
	date, err := time.Parse(LAYOUT, f.date)
	if err != nil {
		return time.Time{}, fmt.Errorf("error in Date(): %w", err)
	}
	return date, nil
}

func (f Filename) Go(d int) Filename {
	date, _ := time.Parse(LAYOUT, f.date)
	newdate := date.AddDate(0, 0, d)
	return NewFilenameFromDate(newdate)
}
