// Copyright 2013 Jesse Allen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shows

import (
	"testing"
	"time"
)

func TestShowAfter(t *testing.T) {
	tests := []struct {
		Name     string
		Expected bool
		A, B     Show
	}{
		{
			"No Dates and No Dates",
			false,
			Show{},
			Show{},
		},
		{
			"No Dates and One Date",
			false,
			Show{},
			Show{Dates: []time.Time{time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"One Date and No Dates",
			true,
			Show{Dates: []time.Time{time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC)}},
			Show{},
		},
		{
			"Many Dates and One Date In the Middle",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"Many Dates Not After One Date",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"Many Dates After One Date",
			true,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"Many Dates Equal to and then After One Date",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"One Date In the Middle of Many Dates",
			true,
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"One Date After Many Dates",
			true,
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"One Date Not After Many Dates",
			false,
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"One Date Equal to but Not After Many Dates",
			false,
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"Many Dates After Many Dates",
			true,
			Show{Dates: []time.Time{
				time.Date(2011, time.June, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.June, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.June, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.June, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"Many Dates Not After Many Dates",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.April, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.April, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.April, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.April, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"Many Dates Not After but Equal to Many Dates",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
	}
	for _, test := range tests {
		if result := test.A.After(test.B); result != test.Expected {
			t.Errorf("%s: Expected: %t Result: %t", test.Name, test.Expected, result)
		}
	}
}

func TestShowBefore(t *testing.T) {
	tests := []struct {
		Name     string
		Expected bool
		A, B     Show
	}{
		{
			"No Dates and No Dates",
			false,
			Show{},
			Show{},
		},
		{
			"No Dates and One Date",
			false,
			Show{},
			Show{Dates: []time.Time{time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"One Date and No Dates",
			true,
			Show{Dates: []time.Time{time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC)}},
			Show{},
		},
		{
			"Many Dates and One Date In the Middle",
			true,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"Many Dates Before One Date",
			true,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"Many Dates Not Before One Date",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"Many Dates Equal to but Not Before One Date",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
		},
		{
			"One Date In the Middle of Many Dates",
			false,
			Show{Dates: []time.Time{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"One Date Not Before Many Dates",
			false,
			Show{Dates: []time.Time{time.Date(2011, time.May, 5, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"One Date Before Many Dates",
			true,
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"One Date Equal to but Not Before Many Dates",
			false,
			Show{Dates: []time.Time{time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"Many Dates Not Before Many Dates",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.June, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.June, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.June, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.June, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"Many Dates Before Many Dates",
			true,
			Show{Dates: []time.Time{
				time.Date(2011, time.April, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.April, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.April, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.April, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
		{
			"Many Dates Not After but Equal to Many Dates",
			false,
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
			Show{Dates: []time.Time{
				time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2011, time.May, 4, 0, 0, 0, 0, time.UTC),
			}},
		},
	}
	for _, test := range tests {
		if result := test.A.Before(test.B); result != test.Expected {
			t.Errorf("%s: Expected: %t Result: %t", test.Name, test.Expected, result)
		}
	}
}
