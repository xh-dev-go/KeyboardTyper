package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/xh-dev-go/KeyboardTyper"
	"os"
)

func main() {
	printMessage := flag.Bool("print", false, "print message instead of typing")
	enter := flag.Bool("enter", false, "enter")
	backspace := flag.Bool("backspace", false, "backspace")
	controlAltDelete := flag.Bool("cad", false, "ctrl+alt+del")
	msg := flag.String("msg", "", "type message")
	fromStd := flag.Bool("stdin", false, "read instructions from stdin")
	flag.Parse()

	if *controlAltDelete == false && *msg == "" {
		fmt.Println("No message to type")
	}

	if fromStd != nil && *fromStd {
		KeyboardTyper.TryTypeFromBuffer(*bufio.NewReader(os.Stdin), *printMessage)
	} else if enter != nil && *enter {
		KeyboardTyper.TryType(KeyboardTyper.InstructionForEnter(), *printMessage)
	} else if backspace != nil && *backspace {
		KeyboardTyper.TryType(KeyboardTyper.InstructionForBackspace(), *printMessage)
	} else if controlAltDelete != nil && *controlAltDelete {
		KeyboardTyper.TryType(KeyboardTyper.InstructionForCAD(), *printMessage)
	} else if msg != nil && *msg != "" {
		KeyboardTyper.TryType(KeyboardTyper.InstructionForString(*msg), *printMessage)
	}
}
