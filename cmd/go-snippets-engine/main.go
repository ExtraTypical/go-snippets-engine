package main

import (
	"fmt"
	"strings"

	hook "github.com/robotn/gohook"
)

func main() {

	fmt.Println("~ Program running! Leader key is " + `";"` + " ~")
	fmt.Println("~ It's recommended that you don't store any " + "information that could be sensitive in nature as a snippet. ~")

	eventHandler()
}

func eventHandler() {

	fmt.Println("~ To close out this program, close the terminal window. ~")

	eventChan := hook.Start()
	defer hook.End()

	readString := ""
	read := false

	for input := range eventChan {

		switch {

		/* Only read if leader is pressed */
		case input.Kind == hook.KeyDown && input.Keychar == ';':
			read = true

		case input.Kind == hook.KeyDown && input.Rawcode == 51:
			if len(readString) > 0 {
				readRune := []rune(readString)
				readString = strings.TrimRight(readString, string(readRune[len(readRune)-1]))
			}
			fmt.Println("String is", readString)

		/* Clear out on space pressed */
		case input.Kind == hook.KeyDown && input.Keychar == ' ':
			readString = ""
			read = false

		/* If read is true, record string */
		case read && input.Kind == hook.KeyDown && input.Keychar != 0:
			readString += string(input.Keychar)
			fmt.Println("String is", readString)
		}

		/* Break on q for now */
		if input.Keychar == 'q' {
			break
		}
	}
}
