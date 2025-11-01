// Package data contians all the structs that are being used for responses
package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   int32     `json:"runtime"`
	Genre     []string  `json:"genere"`
	Version   int32     `json:"version"` // starts from 1 and gets updated `json:""`
	// each time this particular movies is updated
}
