package main

import (
	"KeyboardTyper/KeyboardMapper"
	"bufio"
	"flag"
	"fmt"
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
		KeyboardMapper.TryTypeFromBuffer(*bufio.NewReader(os.Stdin), *printMessage)
	} else if enter != nil && *enter {
		KeyboardMapper.TryType(KeyboardMapper.InstructionForEnter(), *printMessage)
	} else if backspace != nil && *backspace {
		KeyboardMapper.TryType(KeyboardMapper.InstructionForBackspace(), *printMessage)
	} else if controlAltDelete != nil && *controlAltDelete {
		KeyboardMapper.TryType(KeyboardMapper.InstructionForCAD(), *printMessage)
	} else if msg != nil && *msg != "" {
		KeyboardMapper.TryType(KeyboardMapper.InstructionForString(*msg), *printMessage)
	}
}
