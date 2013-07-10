// Copyright 2013 Jesse Allen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shows

import (
	"testing"
	"time"
)

type fakeCalendar []Show

func (f fakeCalendar) Upcoming(limit int, desc bool) []Show {
	return ([]Show)(f)
}

func (f fakeCalendar) Past(limit int, desc bool) []Show {
	return ([]Show)(f)
}

func (f fakeCalendar) equals(g fakeCalendar) bool {
	if len(f) != len(g) {
		return false
	}
	for i, _ := range f {
		if !f[i].minDate().Equal(g[i].minDate()) {
			return false
		}
	}
	return true
}

func TestRegisterCalendar(t *testing.T) {
	c := fakeCalendar{}
	RegisterCalendar(c)
}

func TestUpcoming(t *testing.T) {
	testCals := map[string]fakeCalendar{
		"none": fakeCalendar{},
		"all": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 6, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 7, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 8, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 9, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 10, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 11, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 12, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 13, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 14, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 15, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 16, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 17, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 18, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 19, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 20, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 21, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 22, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 23, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 24, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 25, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 26, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 27, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 28, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 29, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 30, 0, 0, 0, 0, time.UTC)}},
		},
		"even": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 6, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 8, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 10, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 12, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 14, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 16, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 18, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 20, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 22, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 24, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 26, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 28, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 30, 0, 0, 0, 0, time.UTC)}},
		},
		"odd": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 7, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 9, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 11, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 13, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 15, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 17, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 19, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 21, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 23, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 25, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 27, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 29, 0, 0, 0, 0, time.UTC)}},
		},
		"early": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 6, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 7, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 8, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 9, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 10, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 11, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 12, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 13, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 14, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 15, 0, 0, 0, 0, time.UTC)}},
		},
		"late": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 16, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 17, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 18, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 19, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 20, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 21, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 22, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 23, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 24, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 25, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 26, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 27, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 28, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 29, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 30, 0, 0, 0, 0, time.UTC)}},
		},
		"fizzbuzz": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 6, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 9, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 10, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 12, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 15, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 15, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 18, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 20, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 21, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 24, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 25, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 27, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 30, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 30, 0, 0, 0, 0, time.UTC)}},
		},
		"fizz": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 6, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 9, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 12, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 15, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 18, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 21, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 24, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 27, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 30, 0, 0, 0, 0, time.UTC)}},
		},
		"buzz": fakeCalendar{
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 10, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 15, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 20, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 25, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 30, 0, 0, 0, 0, time.UTC)}},
		},
	}
	tests := []struct {
		Expected fakeCalendar
		A, B     fakeCalendar
	}{
		{
			testCals["all"],
			testCals["all"],
			testCals["none"],
		},
		{
			testCals["all"],
			testCals["even"],
			testCals["odd"],
		},
		{
			testCals["all"],
			testCals["odd"],
			testCals["even"],
		},
		{
			testCals["all"],
			testCals["early"],
			testCals["late"],
		},
		{
			testCals["fizzbuzz"],
			testCals["fizz"],
			testCals["buzz"],
		},
	}
	for _, test := range tests {
		// reset source calendars at the start of each test
		calendars = []Calendar{}
		RegisterCalendar(test.A)
		RegisterCalendar(test.B)
		// TODO(jessecarl): update the fake calendars to provide ascending and descending lists
		// TODO(jessecarl): add limit to the tests
		result := Upcoming(0, false)
		if !fakeCalendar(result).equals(test.Expected) {
			t.Errorf("Resulting Calendar does not match Expected Calendar")
		}
	}
}
