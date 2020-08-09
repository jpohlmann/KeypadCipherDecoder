package model

import "strings"

type Keypad struct {
	Grids map[string] Square `json:"square""`
}
type LetterCode struct {
	SquareLocation string `json:"squareLocation"`
	SquarePosition int `json:squarePosition`
}
func GetKeyPad(offset int) Keypad {
	keypad := Keypad{}
	keypad.Grids["UL"] = HydrateSquare(1 + offset % 9)
	keypad.Grids["UC"] = HydrateSquare(2 + offset % 9)
	keypad.Grids["UR"] = HydrateSquare(3 + offset % 9)
	keypad.Grids["CL"] = HydrateSquare(4 + offset % 9)
	keypad.Grids["CC"] = HydrateSquare(5 + offset % 9)
	keypad.Grids["CR"] = HydrateSquare(6 + offset % 9)
	keypad.Grids["DL"] = HydrateSquare(7 + offset % 9)
	keypad.Grids["DC"] = HydrateSquare(8 + offset % 9)
	keypad.Grids["DR"] = HydrateSquare(9 + offset % 9)
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
		for character := 0; character <= len(word); character++ {
			letterCode.SquareLocation += string(word[character])
			if (count == 2) {
				letterCode.SquarePosition = int(word[character])
				decodedMessage += DecodeCharacter(letterCode, decoderPad)
				count = 0
				letterCode.SquareLocation = ""
				letterCode.SquarePosition = 0
				continue
			}
			count++
		}
		decodedMessage += " "
	}
	return decodedMessage
}
