// --Summary:
// A program that displays messages in the terminal using
// other local packages.

package main

import (
	"devguide/demo/pkg/display"
	"devguide/demo/pkg/message"
)

func main() {
  message.Hi()
  display.Display("Hello from the \"display\" package.")
  message.Exciting("An exciting message")
}

