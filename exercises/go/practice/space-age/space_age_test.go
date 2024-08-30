package space

import (
	"math"
	"testing"
)

func TestAge(t *testing.T) {
	const precision = 0.01
	for _, tc := range testCases {
		actual := Age(tc.seconds, tc.planet)
		if math.Abs(actual-tc.expected) > precision {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Age(tc.seconds, tc.planet)
		}
	}
}

var testCases = []struct {
	description string
	seconds     float64
	expected    float64
	planet      Planet
}{
	{
		description: "age on Earth",
		seconds:     1000000000,
		expected:    31.69,
		planet:      "Earth",
	},
	{
		description: "age on Mercury",
		seconds:     2134835688,
		expected:    280.88,
		planet:      "Mercury",
	},
	{
		description: "age on Venus",
		seconds:     189839836,
		expected:    9.78,
		planet:      "Venus",
	},
	{
		description: "age on Mars",
		seconds:     2329871239,
		expected:    39.25,
		planet:      "Mars",
	},
	{
		description: "age on Jupiter",
		seconds:     901876382,
		expected:    2.41,
		planet:      "Jupiter",
	},
	{
		description: "age on Saturn",
		seconds:     3000000000,
		expected:    3.23,
		planet:      "Saturn",
	},
	{
		description: "age on Uranus",
		seconds:     3210123456,
		expected:    1.21,
		planet:      "Uranus",
	},
	{
		description: "age on Neptune",
		seconds:     8210123456,
		expected:    1.58,
		planet:      "Neptune",
	},
}
