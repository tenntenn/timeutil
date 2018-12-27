package timeutil_test

import (
	"fmt"
	"testing"

	"github.com/tenntenn/timeutil"
)

func TestNearHour(t *testing.T) {
	cases := []struct {
		now, expect string
		hour        int
	}{
		{"2018-12-27_00-12-34", "2018-12-27_06-00-00", 6},
		{"2018-12-27_00-12-34", "2018-12-27_18-00-00", 18},
		{"2018-12-27_07-12-34", "2018-12-27_18-00-00", 6},
		{"2018-12-27_07-12-34", "2018-12-27_18-00-00", 18},
		{"2018-12-27_18-12-34", "2018-12-28_06-00-00", 6},
		{"2018-12-31_18-12-34", "2019-01-01_06-00-00", 6},
		{"2018-12-31_18-12-34", "2019-01-01_18-00-00", 18},
		{"2018-12-27_20-30-44", "2018-12-27_21-00-00", 9},
	}

	for _, tt := range cases {
		tt := tt
		n := fmt.Sprintf("%s=%d=>%s", tt.now, tt.hour, tt.expect)
		t.Run(n, func(t *testing.T) {
			now := Parse(t, tt.now)
			expect := Parse(t, tt.expect)
			actual := timeutil.NearHour(now, tt.hour)
			if !actual.Equal(expect) {
				t.Errorf("want %s got %s", tt.expect, Format(t, actual))
			}
		})
	}
}
