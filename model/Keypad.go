package model

import (
	"log"
	"strconv"
	"strings"
)

type Keypad struct {
	Grids map[string] Square `json:"square""`
}
type LetterCode struct {
	SquareLocation string `json:"squareLocation"`
	SquarePosition int `json:squarePosition`
}
func GetKeyPad(offset int) Keypad {
	m := make(map[string]Square)
	keypad := Keypad{m}
	keypad.Grids["UR"] = HydrateSquare((1 + offset) % 9)
	keypad.Grids["UC"] = HydrateSquare((2 + offset) % 9)
	keypad.Grids["UL"] = HydrateSquare((3 + offset) % 9)
	keypad.Grids["SR"] = HydrateSquare((4 + offset) % 9)
	keypad.Grids["SC"] = HydrateSquare((5 + offset) % 9)
	keypad.Grids["SL"] = HydrateSquare((6 + offset) % 9)
	keypad.Grids["DR"] = HydrateSquare((7 + offset) % 9)
	keypad.Grids["DC"] = HydrateSquare((8 + offset) % 9)
	keypad.Grids["DL"] = HydrateSquare((9 + offset) % 9)
	return keypad
}
func DecodeCharacter(code LetterCode, keypad Keypad) string {
	var square Square = keypad.Grids[code.SquareLocation]
	return GetLetter(square, code.SquarePosition)
}
func DecodeMessage(decoderPad Keypad, message string) string {
	wordArray := strings.Split(message, " ")
	decodedMessage := ""
	for arraySpot := 0; arraySpot < len(wordArray); arraySpot++ {
		var count int = 0
		var word string = wordArray[arraySpot]
		letterCode := LetterCode{"", 0}
		for character := 0; character < len(word); character++ {
			if (count == 2) {
				position, err := strconv.Atoi(string(word[character]))
				if err != nil {
					panic(err.Error())
				}
				letterCode.SquarePosition = position
				decodedMessage += DecodeCharacter(letterCode, decoderPad)
				count = 0
				letterCode.SquareLocation = ""
				letterCode.SquarePosition = 0
				continue
			}
			letterCode.SquareLocation += string(word[character])
			count++
		}
		decodedMessage += " "
	}
	return strings.Trim(decodedMessage, " ")
}

func GetBestDecoding(inputString string) string {
	bestDecoding := ""
	score := -1
	for offset := 0; offset < 9; offset++ {
		testDecoderPad := GetKeyPad(offset)
		decodedMessage := DecodeMessage(testDecoderPad, inputString)
		testScore:= ScoreString(decodedMessage)
		log.Print(decodedMessage + " got a score of " + strconv.Itoa(testScore))
		if (testScore > score) {
			bestDecoding = decodedMessage
			score = testScore
		}
	}
	return bestDecoding
}