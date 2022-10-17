package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "must be true",
			args: args{
				str: "madam",
			},
			want: true,
		},
		{
			name: "must be false",
			args: args{
				str: "makan",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.str); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVokalKonsonan(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "must be vokal",
			args: args{
				character: "A",
			},
			want: "Vokal",
		},
		{
			name: "must be Konsonan",
			args: args{
				character: "C",
			},
			want: "Konsonan",
		},
		{
			name: "must be Vokal (lower case)",
			args: args{
				character: "i",
			},
			want: "Vokal",
		},
		{
			name: "must be Konsonal (lower case)",
			args: args{
				character: "p",
			},
			want: "Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVokalKonsonan(tt.args.character); got != tt.want {
				t.Errorf("isVokalKonsonan() = %v, want %v", got, tt.want)
			}
		})
	}
}
