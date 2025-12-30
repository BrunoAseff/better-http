package main

import "testing"

func TestGetGreetingForHour(t *testing.T) {
	tests := []struct {
		hour     int
		expected string
	}{
		{0, "Good night!"},
		{3, "Good night!"},
		{5, "Good night!"},
		{6, "Good morning!"},
		{9, "Good morning!"},
		{11, "Good morning!"},
		{12, "Good afternoon!"},
		{15, "Good afternoon!"},
		{17, "Good afternoon!"},
		{18, "Good evening!"},
		{20, "Good evening!"},
		{23, "Good evening!"},
	}

	for _, tt := range tests {
		result := getGreetingForHour(tt.hour)
		if result != tt.expected {
			t.Errorf("getGreetingForHour(%d) = %s; want %s", tt.hour, result, tt.expected)
		}
	}
}
