package gapbuffer

import "testing"

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
