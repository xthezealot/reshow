package main

import (
	"fmt"
	"testing"
)

func TestCleanName(t *testing.T) {
	dirName = "The Blacklist"
	tests := []struct {
		name string
		want string
	}{
		{"theblacklist-s01e22-final-en-uselessinfo.avi", dirName + " S01E22 FINAL EN.avi"},
		{"theblacklist-s01e22-final-fr-uselessinfo.avi", dirName + " S01E22 FINAL FR.avi"},
		{"theblacklist-s01e01-fr-uselessinfo.avi", dirName + " S01E01 FR.avi"},
		{"theblacklist-s01e01-french-uselessinfo.avi", dirName + " S01E01.avi"},
		{"theblacklist-s01e01-FRENCH-uselessinfo.avi", dirName + " S01E01 FR.avi"},
		{"theblacklist-s01e01-nofr-uselessinfo.avi", dirName + " S01E01.avi"},
		{"theblacklist-s01.avi", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := cleanName(tt.name); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRules(t *testing.T) {
	for re, code := range langsRe {
		fmt.Println(re, code)
	}
}
