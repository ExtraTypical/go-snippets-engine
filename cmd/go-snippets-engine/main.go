package main

import (
	"fmt"

	"github.com/ExtraTypical/go-snippets-engine/internal/handler"
	hook "github.com/robotn/gohook"
)

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

	pressed := handler.PressedKeys{
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
				handler.HandleKeyDown(input, &read, &readString, &pressed)
			}
		case input.Kind == hook.KeyUp:
			{
				handler.HandleKeyUp(input, &pressed)
			}
		case input.Kind == hook.KeyHold:
			{
				handler.HandleKeyHold(input, &pressed)
			}
		}

		/* Review db strings */

	}
}
