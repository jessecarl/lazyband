package shows

import (
	"time"
)

// Show is any engagement at a venue
type Show struct {
	Name  string
	Dates []time.Time // handles single and multi-day shows fluidly
	// TODO(jessecarl): add Poster
	Links        []Link
	OtherArtists []Artist
	Festival     bool
	Headline     bool

	Venue // embed for ease of use (for now)
}

type Venue struct {
	Name string
	Address
	Links []Link
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Artist struct {
	Name  string
	Links []Link
}

type Link struct {
	UID  string
	Name string
	URI  string
}
