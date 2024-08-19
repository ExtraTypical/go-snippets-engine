package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {

	fmt.Println("Program running!")

	eventHandler()
}

func eventHandler() {
	fmt.Println("~ To close out this program, close the terminal window. ~")
	eventChan := hook.Start()
	defer hook.End()

	readString := ""
	read := false

	for input := range eventChan {

		/* Only read if leader is pressed */
		if input.Kind == hook.KeyDown && input.Keychar == ';' {
			read = true
		}

		/* Clear out on space pressed */
		if read && input.Kind == hook.KeyDown && input.Keychar == ' ' {
			readString = ""
			read = false
		}

		/* If read is true, record string */
		if read && input.Kind == hook.KeyDown {
			readString = readString + string(input.Keychar)
			fmt.Println("String is", readString)
		}

		/* Break on q for now */
		if input.Keychar == 'q' {
			break
		}
	}

}
