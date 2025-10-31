// Package data contians all the structs that are being used for responses
package data

import "time"

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   int32
	Genre     []string
	Version   int32 // starts from 1 and gets updated
	// each time this particular movies is updated
}
