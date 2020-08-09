package model

import "log"

type Square struct {
	PositionOne string `json:"positionOne"`
	PositionTwo string `json:"positionTwo"`
	PositionThree string `json:"positionThree"`
}
func GetLetter(square Square, number int) string {
	switch number {
	case 1:
		return square.PositionOne
	case 2:
		return square.PositionTwo
	case 3:
		return square.PositionThree
	default:
		return ""
	}
}
func HydrateSquare(number int) Square {
	hydratedSquare := Square{}
	switch number {
	case 1:
		hydratedSquare.PositionOne="A"
		hydratedSquare.PositionTwo="B"
		hydratedSquare.PositionThree="C"
	case 2:
		hydratedSquare.PositionOne="D"
		hydratedSquare.PositionTwo="E"
		hydratedSquare.PositionThree="F"
	case 3:
		hydratedSquare.PositionOne="G"
		hydratedSquare.PositionTwo="H"
		hydratedSquare.PositionThree="I"
	case 4:
		hydratedSquare.PositionOne="J"
		hydratedSquare.PositionTwo="K"
		hydratedSquare.PositionThree="L"
	case 5:
		hydratedSquare.PositionOne="M"
		hydratedSquare.PositionTwo="N"
		hydratedSquare.PositionThree="O"
	case 6:
		hydratedSquare.PositionOne="P"
		hydratedSquare.PositionTwo="Q"
		hydratedSquare.PositionThree="R"
	case 7:
		hydratedSquare.PositionOne="S"
		hydratedSquare.PositionTwo="T"
		hydratedSquare.PositionThree="U"
	case 8:
		hydratedSquare.PositionOne="V"
		hydratedSquare.PositionTwo="W"
		hydratedSquare.PositionThree="X"
	case 0:
		hydratedSquare.PositionOne="Y"
		hydratedSquare.PositionTwo="Z"
	default:
		log.Fatal("Number outside of range")
	}
	return hydratedSquare
}