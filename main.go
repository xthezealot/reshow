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
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// S01E01 match
	seRe = regexp.MustCompile(`(?i)S\d\dE\d\d`)
	// Final match
	finRe = regexp.MustCompile(`(?i)FINAL`)
	// Language match
	langsRe = map[*regexp.Regexp]string{
		regexp.MustCompile(`(?i)[^A-Z]EN[^A-Z]|ENGLISH`): "EN",
		regexp.MustCompile(`(?i)[^A-Z]FR[^A-Z]|FRENCH`):  "FR",
	}
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fileName := f.Name()
		if !seRe.MatchString(fileName) {
			continue
		}
		newName := currentDirName()
		newName += " " + strings.ToUpper(seRe.FindString(fileName))
		if finRe.MatchString(fileName) {
			newName += " " + "FINAL"
		}
		for re, lng := range langsRe {
			if re.MatchString(fileName) {
				newName += " " + lng
			}
		}
		newName += filepath.Ext(fileName)
		if err = os.Rename(fileName, newName); err != nil {
			panic(err)
		}
	}
}

func currentDirName() string {
	dir, _ := os.Getwd()
	return filepath.Base(dir)
}
