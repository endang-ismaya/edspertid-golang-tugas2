package main

import "testing"

func Test_palindromeMenu(t *testing.T) {
	tests := []struct {
		name string
	}{
		name: "madam",
		want: true
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			palindromeMenu()
		})
	}
}
