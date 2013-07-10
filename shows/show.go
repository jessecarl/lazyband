// Copyright 2013 Jesse Allen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// Returns true if the Show s begins after the Show t begins.
// Shows with no dates will evaluate as not after shows with dates.
func (s Show) After(t Show) bool {
	sDate := s.minDate()
	tDate := t.minDate()
	if sDate.IsZero() {
		return false
	} else if tDate.IsZero() {
		return true
	}
	return sDate.After(tDate)
}

// Returns true if Show s begins before Show t begins.
// Shows with no dates will evaluate as not before shows with dates.
// This assumes that the use of Before is a positive assertion.
func (s Show) Before(t Show) bool {
	sDate := s.minDate()
	tDate := t.minDate()
	if sDate.IsZero() {
		return false
	} else if tDate.IsZero() {
		return true
	}
	return sDate.Before(tDate)
}

func (s Show) minDate() time.Time {
	if len(s.Dates) < 1 {
		return time.Time{} // give a zero-valued time if no dates (an unexpected case)
	}
	t := s.Dates[0]
	if len(s.Dates) == 1 {
		return t
	}
	for _, u := range s.Dates[1:] {
		if t.After(u) {
			t = u
		}
	}
	return t
}
