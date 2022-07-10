package main

import (
	"testing"
)

func TestEnlarge(t *testing.T) {
	case1 := rect{
		width:  2,
		height: 2,
	}

	t.Log("Given rect:", case1)

	t.Run("wider", func(t *testing.T) {
		t.Log("When make it width twice larger")

		case1.makeWider(2)

		t.Log("Then it should be increased:", case1)
		if case1.width != 4 {
			t.Error("Width should be enlarged")
		}
	})

	t.Run("higher", func(t *testing.T) {
		t.Log("When make it height twice larger")

		case1.makeHigher(2)

		t.Log("Then it should be increased:", case1)
		if case1.height != 4 {
			t.Error("Height should be enlarged")
		}
	})

	t.Run("both", func(t *testing.T) {
		if case1.width != 4 && case1.height != 4 {
			t.Error("Changes should be preserved")
		}
	})
}
