package gigasecond

import (
	"os"
	"testing"
	"time"
)

// Load the test cases from a JSON file
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var testCases = []struct {
	description string
	in          string
	want        string
}{
	{
		description: "date only specification of time",
		in:          "2011-04-25",
		want:        "2043-01-01T01:46:40Z",
	},
	{
		description: "second test for date only specification of time",
		in:          "1977-06-13",
		want:        "2009-02-19T01:46:40Z",
	},
	{
		description: "third test for date only specification of time",
		in:          "1959-07-19",
		want:        "1991-03-27T01:46:40Z",
	},
	{
		description: "full time specified",
		in:          "2015-01-24T22:00:00Z",
		want:        "2046-10-02T23:46:40Z",
	},
	{
		description: "full time with day roll-over",
		in:          "2015-01-24T23:59:59Z",
		want:        "2046-10-03T01:46:39Z",
	},
}

func TestAddGigasecond(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			in := parse(tc.in, t)
			want := parse(tc.want, t)
			got := AddGigasecond(in)
			if !got.Equal(want) {
				t.Errorf("AddGigasecond(%v) = %v, want: %v", in, got, want)
			}
		})
	}
}

func parse(s string, t *testing.T) time.Time {
	tt, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t.Fatalf("error in test setup: invalid RFC3339Nano time string: %v", err)
	}
	return tt
}

func BenchmarkAddGigasecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddGigasecond(time.Date(2011, time.April, 25, 0, 0, 0, 0, time.UTC))
	}
}
