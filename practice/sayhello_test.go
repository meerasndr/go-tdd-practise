package main

import (
	"testing"
)

var languagetests = []struct{
	language	string
	want 		string
}{
	{"Spanish", "Hola!"},
	{"French", "Bonjour!"},
	{"Sanskrit", "Namaskaram"},
	{"Tamil", "Vanakkam"},
}

func TestHello(t *testing.T){
	for _, tt := range languagetests{
		t.Run(tt.language, func(t *testing.T){
			got := Hello(tt.language)
			if got != tt.want {
				t.Errorf("Got %q, want %q", got, tt.want)
			}
		})
	}
}