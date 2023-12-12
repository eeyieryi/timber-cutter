package game

import (
	"math/rand"
)

type Position int

const (
	PosLeft  Position = -1
	PosNone  Position = 0 // Empty
	PosRight Position = 1
)

type GameState struct {
	PlayerPos Position
	Tree      []Position
}

func NewGameState() *GameState {
	return &GameState{
		PlayerPos: PosLeft,
		Tree:      GenerateInitialTree(),
	}
}

func GenerateInitialTree() []Position {
	sliceLength := 10
	sections := make([]Position, sliceLength)
	sections[0] = PosNone // initial tree always start with the root empty
	for i := 1; i < sliceLength; i++ {
		sections[i] = GenerateSection(sections[i-1])
	}
	return sections
}

func GenerateSection(prevPos Position) Position {
	var possiblePos = []Position{PosNone} // can always add empty

	switch prevPos {
	case PosNone:
		// can only add *any* position or empty
		possiblePos = append(possiblePos, PosLeft, PosRight)
	case PosLeft, PosRight:
		// can only add the *same* position or empty
		possiblePos = append(possiblePos, prevPos)
	}

	rIndex := rand.Intn(len(possiblePos))
	return possiblePos[rIndex]
}

func (g *GameState) UpdateTree() {
	lastIndex := len(g.Tree) - 1
	newSection := GenerateSection(g.Tree[lastIndex])
	g.Tree = append(g.Tree[1:], newSection)
}
