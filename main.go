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
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
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
	langs   = []language.Tag{
		language.Albanian,
		language.Armenian,
		language.Bulgarian,
		language.Catalan,
		language.Chinese,
		language.Croatian,
		language.Czech,
		language.Danish,
		language.Dutch,
		language.English,
		language.Estonian,
		language.Filipino,
		language.Finnish,
		language.French,
		language.Georgian,
		language.Greek,
		language.Hebrew,
		language.Hindi,
		language.Hungarian,
		language.Icelandic,
		language.Indonesian,
		language.Italian,
		language.Japanese,
		language.Kazakh,
		language.Korean,
		language.Latvian,
		language.Lithuanian,
		language.Macedonian,
		language.Malay,
		language.Mongolian,
		language.Nepali,
		language.Norwegian,
		language.Persian,
		language.Polish,
		language.Portuguese,
		language.Romanian,
		language.Russian,
		language.Serbian,
		language.Slovak,
		language.Slovenian,
		language.Spanish,
		language.Swahili,
		language.Swedish,
		language.Thai,
		language.Turkish,
		language.Ukrainian,
		language.Uzbek,
		language.Vietnamese,
	}
)

func init() {
	langNamerEN := display.Languages(language.English)
	for _, l := range langs {
		re := regexp.MustCompile("((?i)[^A-Z]" + l.String() + "[^A-Z])|([^A-Z]" + strings.ToUpper(langNamerEN.Name(l)) + "[^A-Z])|([^A-Z]" + strings.ToUpper(display.Self.Name(l)) + "[^A-Z])")
		langsRe[re] = strings.ToUpper(l.String())
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
