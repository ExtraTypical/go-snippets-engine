package main

import (
	"fmt"
	"strings"

	"github.com/ExtraTypical/go-snippets-engine/internal/window"
	hook "github.com/robotn/gohook"
)

type PressedKeys struct {
	Ctrl  bool
	Shift bool
	O     bool
}

func main() {

	fmt.Println("~ Program running! Leader key is " + `";"` + " ~")
	fmt.Println("~ Don't store any " + "information that could be sensitive in nature as a snippet. ~")
	fmt.Println("~ Press Ctrl+Shift+O to open a window for adding new snippets ~")

	eventHandler()
}

func eventHandler() {

	fmt.Println("~ To close out this program, close the terminal window. ~")

	eventChan := hook.Start()
	defer hook.End()

	readString := ""
	read := false

	pressed := PressedKeys{
		Ctrl:  false,
		Shift: false,
		O:     false,
	}

	for input := range eventChan {

		// if input.Keychar != 0 {
		// 	fmt.Println(input)
		// }

		/* Updated refactored method */
		switch {
		case input.Kind == hook.KeyDown:
			{
				handleKeyDown(input, &read, &readString, &pressed)
			}
		case input.Kind == hook.KeyUp:
			{
				handleKeyUp(input, &pressed)
			}
		case input.Kind == hook.KeyHold:
			{
				handleKeyHold(input, &pressed)
			}
		}

		/* Review db strings */

	}
}

func handleKeyDown(input hook.Event, read *bool, readString *string, pressed *PressedKeys) {

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

func handleKeyUp(input hook.Event, pressed *PressedKeys) {

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

func handleKeyHold(input hook.Event, pressed *PressedKeys) {
	if input.Rawcode == 59 {
		pressed.Ctrl = true
	}
	if input.Rawcode == 56 {
		pressed.Shift = true
	}
}
