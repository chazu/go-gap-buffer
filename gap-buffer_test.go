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

	if x != 4 || y != 1 {
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
	g.MoveCursorUp()
	if g.GetCursorIndex() != 5 {
		t.Errorf("Expected cursor index of 5, got %d", g.GetCursorIndex())
	}
}
