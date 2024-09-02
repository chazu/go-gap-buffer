package gapbuffer

import (
	"testing"
)

func TestGapBufferSetString(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{
			name: "empty string",
			s:    "",
		},
		{
			name: "non-empty string",
			s:    "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := new(GapBuffer)
			g.SetString(tt.s)
			if g.GetString() != tt.s {
				t.Errorf("expected %s, got %s", tt.s, g.GetString())
			}
		})
	}
}

func TestGapBufferInsertWhenEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panic occurred when inserting into empty buffer")
		}
	}()

	g := new(GapBuffer)
	g.SetString("")
	expected := 'a'
	g.Insert('a')

	if g.GetString() != "a" {
		t.Errorf("expected %s, got %s", string(expected), g.GetString())
	}
}

func TestGetCursorXYEmptyBuffer(t *testing.T) {
	g := new(GapBuffer)
	x, y := g.GetCursorXY()
	if x != 0 || y != 0 {
		t.Errorf("expected 0, 0, got %d, %d", x, y)
	}
}

// When we set the string, the cursor should be at the beginning of the string
func TestGetCursorXYNonEmptyBuffer(t *testing.T) {
	g := new(GapBuffer)
	g.SetString("hello")
	x, y := g.GetCursorXY()
	if x != 0 || y != 0 {
		t.Errorf("expected 0, 0, got %d, %d", x, y)
	}
}

// When we insert runes one by one, the cursor should be at the end
func TestGetCursorXYBufferWithTextAdded(t *testing.T) {
	g := new(GapBuffer)

	text := "01234\n0123"
	for _, r := range text {
		g.Insert(r)
	}
	x, y := g.GetCursorXY()

	if x != 4 && y != 1 {
		t.Errorf("expected 4, 1, got %d, %d", x, y)
	}
}

func TestMoveCursorLeftOne(t *testing.T) {
	g := new(GapBuffer)
	text := "hello"
	for _, r := range text {
		g.Insert(r)
	}
	g.MoveCursorLeft(1)
	if g.GetCursorIndex() != 4 {
		t.Errorf("Expected cursor index of 4, got %d", g.GetCursorIndex())
	}
}

func TestMoveCursorLeftTwo(t *testing.T) {
	g := new(GapBuffer)
	text := "hello"
	for _, r := range text {
		g.Insert(r)
	}
	g.MoveCursorLeft(2)
	if g.GetCursorIndex() != 3 {
		t.Errorf("Expected cursor index of 3, got %d", g.GetCursorIndex())
	}
}

func TestMoveCursorUp(t *testing.T) {
	g := new(GapBuffer)
	text := "hello\nworld"
	for _, r := range text {
		g.Insert(r)
	}
	g.MoveCursorUp(1)
	if g.GetCursorIndex() != 5 {
		t.Errorf("Expected cursor index of 5, got %d", g.GetCursorIndex())
	}
}

func TestLines(t *testing.T) {
	g := new(GapBuffer)
	text := "hello\nworld"
	for _, r := range text {
		g.Insert(r)
	}
	lines := g.lines()
	if len(lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(lines))
	}
	if lines[0] != "hello" {
		t.Errorf("Expected 'hello', got %s", lines[0])
	}
	if lines[1] != "world" {
		t.Errorf("Expected 'world', got %s", lines[1])
	}
}

func TestDistanceToMoveBack(t *testing.T) {
	tests := []struct {
		name         string
		x            int
		y            int
		previousLine string
		expected     int
	}{
		{
			name:         "empty previous line",
			x:            0,
			y:            0,
			previousLine: "",
			expected:     1,
		},
		{
			name:         "non-empty previous line",
			x:            2,
			y:            1,
			previousLine: "hello",
			expected:     6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if distance := distanceToMoveBack(tt.x, tt.y, tt.previousLine); distance != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, distance)
			}
		})
	}
}

func TestDistanceToMoveForward(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        int
		thisLine string
		expected int
	}{
		{
			name:     "empty this line",
			x:        0,
			y:        0,
			thisLine: "",
			expected: 1,
		},
		{
			name:     "non-empty this line",
			x:        2,
			y:        1,
			thisLine: "hello",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if distance := distanceToMoveForward(tt.x, tt.y, tt.thisLine); distance != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, distance)
			}
		})
	}
}
