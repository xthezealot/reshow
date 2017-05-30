/*
Reshow cleans file names in the current directory that refer to a TV show (with season and episode).

The current directory is supposed to represent the show title.
For example, if you are in the "The Blacklist" directory, it renames the "theblacklist-s01e22-final-fr-uselessinfo.avi" file into "The Blacklist S01E22 FINAL FR.avi".

Usage

1. Go to the show directory.
	cd ~/Movies/The\ Blacklist

2. Run Reshow.
	reshow
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	dirName string
	seRe    = regexp.MustCompile(`(?i)S\d\dE\d\d`) // S01E01 match
	finalRe = regexp.MustCompile(`(?i)FINAL`)
	langsRe = make(map[*regexp.Regexp]string)
	langs   = map[string]string{
		"EN": "ENGLISH",
		"DE": "GERMAN",
		"FR": "FRENCH",
		"NL": "DUTCH",
		"RO": "ROMANIAN",
		"JP": "JAPANESE",
		"RU": "RUSSIAN",
	}
)

func init() {
	for code, l := range langs {
		langsRe[regexp.MustCompile("((?i)[^A-Z]"+code+"[^A-Z])|([^A-Z]"+l+"[^A-Z])")] = code
	}
}

func main() {
	var err error
	dirName, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	dirName = filepath.Base(dirName)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		n, err := cleanName(f.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err = os.Rename(f.Name(), n); err != nil {
			panic(err)
		}
	}
}

func cleanName(file string) (n string, err error) {
	if !seRe.MatchString(file) {
		err = fmt.Errorf("%s is not a TV show name", file)
		return
	}
	n = dirName + " " + strings.ToUpper(seRe.FindString(file))
	if finalRe.MatchString(file) {
		n += " " + "FINAL"
	}
	for re, lng := range langsRe {
		if re.MatchString(file) {
			n += " " + lng
		}
	}
	n += filepath.Ext(file)
	return
}
