package gigasecond

import (
	"os"
	"testing"
	"time"
)

// Test_AddGigasecond tests the AddGigasecond function.
func TestAddGigasecond(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"date only specification of time",
			"2011-04-25",
			"2043-01-01T01:46:40Z",
		},
		{
			"second test for date only specification of time",
			"1977-06-13",
			"2009-02-19T01:46:40Z",
		},
		{
			"third test for date only specification of time",
			"1959-07-19",
			"1991-03-27T01:46:40Z",
		},
		{
			"full time specified",
			"2015-01-24T22:00:00Z",
			"2046-10-02T23:46:40Z",
		},
		{
			"full time with day roll-over",
			"2015-01-24T23:59:59Z",
			"2046-10-03T01:46:39Z",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var input time.Time
			var err error
			if len(tc.input) == 10 {
				input, err = time.Parse("2006-01-02", tc.input)
			} else {
				input, err = time.Parse(time.RFC3339, tc.input)
			}
			if err != nil {
				t.Fatalf("Failed to parse input: %v", err)
			}
			want, err := time.Parse(time.RFC3339Nano, tc.expected)
			if err != nil {
				t.Fatalf("Failed to parse expected: %v", err)
			}
			got := AddGigasecond(input)
			if !got.Equal(want) {
				t.Errorf("AddGigasecond(%v) = %v, want %v", input, got, want)
			}
		})
	}
}

func BenchmarkAddGigasecond(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		AddGigasecond(time.Now())
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}