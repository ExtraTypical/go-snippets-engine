package handler

import (
	"fmt"
	"strings"

	"github.com/ExtraTypical/go-snippets-engine/internal/window"
	hook "github.com/robotn/gohook"
)

func HandleKeyDown(input hook.Event, read *bool, readString *string, pressed *PressedKeys) {

	if input.Rawcode == 31 {
		pressed.O = true
	}

	if pressed.Ctrl && pressed.Shift && pressed.O {
		window.OpenWindow()

	}

	switch {
	/* Only read if leader is pressed */
	case input.Kind == hook.KeyDown && input.Keychar == ';':
		*read = true

	/* On backspace, stop recording */
	case input.Kind == hook.KeyDown && input.Rawcode == 51:
		if len(*readString) > 0 {
			readRune := []rune(*readString)
			*readString = strings.TrimRight(*readString, string(readRune[len(readRune)-1]))
			fmt.Println("String is", *readString)
		} else {
			*read = false
		}

	/* Clear out on space pressed */
	case input.Kind == hook.KeyDown && input.Keychar == ' ':
		*readString = ""
		*read = false

	/* If read is true, record string */
	case *read && input.Kind == hook.KeyDown && input.Keychar != 0:
		*readString += string(input.Keychar)
		fmt.Println("String is", *readString)

	}
}

func HandleKeyUp(input hook.Event, pressed *PressedKeys) {

	if input.Rawcode == 59 {
		pressed.Ctrl = false
	}
	if input.Rawcode == 56 {
		pressed.Shift = false
	}
	if input.Rawcode == 31 {
		pressed.O = false
	}
}

func HandleKeyHold(input hook.Event, pressed *PressedKeys) {
	if input.Rawcode == 59 {
		pressed.Ctrl = true
	}
	if input.Rawcode == 56 {
		pressed.Shift = true
	}
}
