// Package provides database interactions and manipulation of objects
package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"strings"
)

// Our database connector object
type DBManager struct {
	db            *gorm.DB
	isInitialized bool
}

// Let's make a singleton instance of the database connection
var dbInstance = connect()

// Our word object.  We alphabetize the letters in the word to make it easy
// to make an indexed search for anagrams.
type Word struct {
	AlphabetizedChars string `gorm:"size:50;index:alpha;primary_key"`
	Word              string `gorm:"size:50;primary_key"`
}

// Connect to our sqlite database
func connect() DBManager {
	localDbRef, err := gorm.Open("sqlite3", "./dictionary.db")
	localDbRef.AutoMigrate(&Word{})
	if err != nil {
		panic(err.Error())
	} else {
		log.Print("DB Initialized successfully")
	}
	return DBManager{db: localDbRef, isInitialized: true}
}

// Get the metadata for the words
func LookupWord(wordToLookup string) bool {
	foundWords := make([]Word, 0)
	dbInstance.db.Where("SELECT * FROM words WHERE word = ?", wordToLookup).First(&foundWords)
	return len(foundWords) > 0
}

func ScoreString(resultString string) int {
	score := 0
	wordArray := strings.Split(resultString, " ")
	for arraySpot := 0; arraySpot < len(wordArray); arraySpot++ {
		if (LookupWord(wordArray[arraySpot])) {
			score++
		}
	}
	return score
}
