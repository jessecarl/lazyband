// Copyright 2013 Jesse Allen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shows

import (
	"testing"
	"time"
)

// Assume that all fakeCalendars are set up with shows in ascending order
type fakeCalendar []Show

func (f fakeCalendar) Upcoming(limit int, desc bool) []Show {
	s := ([]Show)(f)
	if desc {
		t := []Show{}
		for i := len(s); i > 0; i-- {
			t = append(t, s[i-1])
		}
		s = t
	}
	if limit > 0 && len(s) > limit {
		s = s[:limit]
	}
	return s
}

func (f fakeCalendar) Past(limit int, desc bool) []Show {
	s := ([]Show)(f)
	if desc {
		t := []Show{}
		for i := len(s); i > 0; i-- {
			t = append(t, s[i-1])
		}
		s = t
	}
	if limit > 0 && len(s) > limit {
		s = s[:limit]
	}
	return s
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

var testCals map[string]fakeCalendar

func init() {
	testCals = map[string]fakeCalendar{
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
}

func TestRegisterCalendar(t *testing.T) {
	c := fakeCalendar{}
	RegisterCalendar(c)
}

func TestUpcoming(t *testing.T) {
	tests := []struct {
		Expected, A, B fakeCalendar
		Limit          int
		Desc           bool
	}{
		{
			testCals["all"],
			testCals["all"],
			testCals["none"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["even"],
			testCals["odd"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["odd"],
			testCals["even"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["early"],
			testCals["late"],
			30,
			false,
		},
		{
			testCals["fizzbuzz"],
			testCals["fizz"],
			testCals["buzz"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["all"],
			testCals["none"],
			15,
			false,
		},
		{
			testCals["all"],
			testCals["even"],
			testCals["odd"],
			20,
			false,
		},
		{
			testCals["all"],
			testCals["odd"],
			testCals["even"],
			10,
			false,
		},
		{
			testCals["all"],
			testCals["early"],
			testCals["late"],
			15,
			false,
		},
		{
			testCals["fizzbuzz"],
			testCals["fizz"],
			testCals["buzz"],
			45,
			false,
		},
		{
			testCals["all"],
			testCals["all"],
			testCals["none"],
			30,
			true,
		},
		{
			testCals["all"],
			testCals["even"],
			testCals["odd"],
			20,
			true,
		},
		{
			testCals["all"],
			testCals["odd"],
			testCals["even"],
			30,
			true,
		},
		{
			testCals["all"],
			testCals["early"],
			testCals["late"],
			15,
			true,
		},
		{
			testCals["fizzbuzz"],
			testCals["fizz"],
			testCals["buzz"],
			45,
			true,
		},
	}
	for _, test := range tests {
		// reset source calendars at the start of each test
		calendars = []Calendar{}
		RegisterCalendar(test.A)
		RegisterCalendar(test.B)
		result := Upcoming(test.Limit, test.Desc)
		if !fakeCalendar(result).equals(fakeCalendar(test.Expected.Upcoming(test.Limit, test.Desc))) {
			t.Errorf("Resulting Calendar does not match Expected Calendar")
		}
	}
}

func TestPast(t *testing.T) {
	tests := []struct {
		Expected, A, B fakeCalendar
		Limit          int
		Desc           bool
	}{
		{
			testCals["all"],
			testCals["all"],
			testCals["none"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["even"],
			testCals["odd"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["odd"],
			testCals["even"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["early"],
			testCals["late"],
			30,
			false,
		},
		{
			testCals["fizzbuzz"],
			testCals["fizz"],
			testCals["buzz"],
			30,
			false,
		},
		{
			testCals["all"],
			testCals["all"],
			testCals["none"],
			15,
			false,
		},
		{
			testCals["all"],
			testCals["even"],
			testCals["odd"],
			20,
			false,
		},
		{
			testCals["all"],
			testCals["odd"],
			testCals["even"],
			10,
			false,
		},
		{
			testCals["all"],
			testCals["early"],
			testCals["late"],
			15,
			false,
		},
		{
			testCals["fizzbuzz"],
			testCals["fizz"],
			testCals["buzz"],
			45,
			false,
		},
		{
			testCals["all"],
			testCals["all"],
			testCals["none"],
			30,
			true,
		},
		{
			testCals["all"],
			testCals["even"],
			testCals["odd"],
			20,
			true,
		},
		{
			testCals["all"],
			testCals["odd"],
			testCals["even"],
			30,
			true,
		},
		{
			testCals["all"],
			testCals["early"],
			testCals["late"],
			15,
			true,
		},
		{
			testCals["fizzbuzz"],
			testCals["fizz"],
			testCals["buzz"],
			45,
			true,
		},
	}
	for _, test := range tests {
		// reset source calendars at the start of each test
		calendars = []Calendar{}
		RegisterCalendar(test.A)
		RegisterCalendar(test.B)
		result := Past(test.Limit, test.Desc)
		if !fakeCalendar(result).equals(fakeCalendar(test.Expected.Past(test.Limit, test.Desc))) {
			t.Errorf("Resulting Calendar does not match Expected Calendar")
		}
	}
}
