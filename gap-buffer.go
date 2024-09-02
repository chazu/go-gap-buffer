package gapbuffer

import (
	"strings"
)

const NEWLINE = '\n'

type GapBuffer struct {
	buffer     []rune
	preGapLen  int
	postGapLen int
}

// gapStart returns the index at which the gap is starting
func (g *GapBuffer) gapStart() int {

	return g.preGapLen
}

// gapLen returns the length of the gap
func (g *GapBuffer) gapLen() int {

	return g.postGapStart() - g.preGapLen
}

// postGapStart returns the immediately next index after the end of the gap index
func (g *GapBuffer) postGapStart() int {

	return len(g.buffer) - g.postGapLen
}

// Get the string as lines, split on newlines
func (g *GapBuffer) lines() []string {
	return strings.Split(g.GetString(), string(NEWLINE))
}

func (g *GapBuffer) GetCursorIndex() int {
	return g.preGapLen
}

func (g *GapBuffer) GetCursorXY() (int, int) {
	lines := g.lines()
	cursorIndex := g.GetCursorIndex()
	line := 0
	for i, l := range lines {
		if cursorIndex <= len(l) {

			line = i
			break
		}
		// Remove extra 1 for newline character
		cursorIndex -= len(l) + 1
	}
	return cursorIndex, line
}

// SetString initialises the buffer with the characters of the input string
func (g *GapBuffer) SetString(s string) {

	g.buffer = []rune(s)

	// initialise preGap and postGap according to the input string
	g.preGapLen = 0
	g.postGapLen = len(g.buffer)
}

// GetString returns the text stored in the buffer
func (g *GapBuffer) GetString() string {

	// create a new rune slice and append the preGap and postGap slices to it before returning
	text := append([]rune{}, g.buffer[:g.preGapLen]...)
	text = append(text, g.buffer[g.postGapStart():]...)

	return string(text)
}

// MoveCursorRight moves the cursor position to the right by one step
func (g *GapBuffer) MoveCursorRight() {

	// check if the cursor is at the end of the buffer
	if g.postGapLen == 0 {
		return
	}

	// copy the elements from the rear to the front of the gap to shift the gap towards right
	g.buffer[g.preGapLen] = g.buffer[g.postGapStart()]
	g.preGapLen++
	g.postGapLen--
}

// MoveCursorLeft moves the cursor position to the left by one step
func (g *GapBuffer) MoveCursorLeft(distance int) {

	// check if the cursor is at the start of the buffer
	if g.preGapLen == 0 {
		return
	}

	// copy the elements from the front to the rear of the gap to shift the gap towards left
	g.buffer[g.postGapStart()-distance] = g.buffer[g.preGapLen-distance]
	g.preGapLen -= distance
	g.postGapLen += distance
}

// Calculate the distance to move back when moving up a line
func distanceToMoveBack(x int, y int, previousLine string) int {

	// Chars on previous line after the current column
	previousLineChars := previousLine[x:]

	return x + len(previousLineChars) + 1
}

// Calculate the distance to move forward when moving up a line
func distanceToMoveForward(x int, y int, thisLine string) int {

	// Chars on this line after the current column
	thisLineChars := thisLine[x:]

	return x + len(thisLineChars) + 1
}

func (g *GapBuffer) MoveCursorUp(distance int) {
	x, y := g.GetCursorXY()

	for i := 0; i < distance; i += 1 {
		// Get characters after cursor on previous line
		previousLine := g.lines()[y-1]

		// Calculate how far back to move the cursor index - add one for the newline character
		moveBack := distanceToMoveBack(x, y, previousLine)
		g.MoveCursorLeft(moveBack)

		// decrement y for next go-round
		y--
	}
}

func (g *GapBuffer) MoveCursorDown(distance int) {
	x, y := g.GetCursorXY()

	for i := 0; i < distance; i += 1 {
		// Calculate how far back to move the cursor index - add one for the newline character
		moveBack := distanceToMoveBack(x, y, g.lines()[y])
		g.MoveCursorLeft(moveBack)

		// decrement y for next go-round
		y--
	}
}

// Delete deletes a character immediately after the cursor
func (g *GapBuffer) Delete() {

	// check if the cursor is at the end of the buffer
	if g.postGapLen == 0 {
		return
	}

	// shrink postGap from the start
	g.postGapLen--
}

// Backspace deletes a character immediately before the cursor
func (g *GapBuffer) Backspace() {

	// check if the cursor is at the start of the buffer
	if g.preGapLen == 0 {
		return
	}

	// shrink preGap from the end
	g.preGapLen--
}

// growGap creates a gap of length equal to the buffer length to be created between preGap and postGap
func (g *GapBuffer) growGap() {

	toGrow := len(g.buffer)
	if toGrow == 0 {
		toGrow = 1
	}

	// create a new rune slice of length equal to twice the buffer length and copy the
	// preGap elements and the postGap elements in it such that a gap of length equal
	// to the buffer length is created between them before assigning it as the buffer
	newBuffer := make([]rune, toGrow*2)

	copy(newBuffer, g.buffer[:g.preGapLen])
	copy(newBuffer[g.postGapStart()+len(g.buffer):], g.buffer[g.postGapStart():])

	g.buffer = newBuffer
}

// Insert inserts a single character at the cursor position
func (g *GapBuffer) Insert(c rune) {

	// grow the gap if necessary so that insertion can take place at the start of the gap
	if g.gapLen() == 0 {
		g.growGap()
	}

	g.buffer[g.gapStart()] = c
	g.preGapLen++
}
