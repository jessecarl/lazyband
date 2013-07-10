// Copyright 2013 Jesse Allen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shows

type Calendar interface {
	Upcoming(limit int, desc bool) []Show
	Past(limit int, desc bool) []Show
}

var calendars []Calendar

// Registers calendars with the global calendar to be queried
// together.
func RegisterCalendar(c Calendar) {
	calendars = append(calendars, c)
}

// Sorts and merges results from all calendars
func Upcoming(limit int, desc bool) []Show {
	s := []Show{}
	for _, c := range calendars {
		s = mergeShows(s, c.Upcoming(limit, desc), (func(d bool) func(Show, Show) bool {
			if d {
				return func(a, b Show) bool { return a.After(b) }
			} else {
				return func(a, b Show) bool { return a.Before(b) }
			}
		}(desc)))
	}
	if limit > 0 && len(s) > limit {
		s = s[:limit]
	}
	return s
}

// Assumes that input slices are already sorted
func mergeShows(a, b []Show, compare func(Show, Show) bool) []Show {
	c := []Show{}
	for len(a) > 0 && len(b) > 0 {
		if compare(a[0], b[0]) {
			c = append(c, a[0])
			if len(a) > 1 {
				a = a[1:]
			} else {
				a = []Show{}
			}
		} else {
			c = append(c, b[0])
			if len(b) > 1 {
				b = b[1:]
			} else {
				b = []Show{}
			}
		}
	}
	if len(a) > 0 {
		c = append(c, a...)
	}
	if len(b) > 0 {
		c = append(c, b...)
	}
	return c
}
