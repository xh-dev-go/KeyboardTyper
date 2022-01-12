package KeyboardTyper

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"
)

const KEY_MOD_LCTRL = 0x01
const KEY_MOD_LSHIFT = 0x02
const KEY_MOD_LALT = 0x04
const KEY_MOD_LMETA = 0x08
const KEY_MOD_RCTRL = 0x10
const KEY_MOD_RSHIFT = 0x20
const KEY_MOD_RALT = 0x40
const KEY_MOD_RMETA = 0x80

//const KEY_A = 0x04 // Keyboard a and A
//const KEY_B = 0x05 // Keyboard b and B
//const KEY_C = 0x06 // Keyboard c and C
//const KEY_D = 0x07 // Keyboard d and D
//const KEY_E = 0x08 // Keyboard e and E
//const KEY_F = 0x09 // Keyboard f and F
//const KEY_G = 0x0a // Keyboard g and G
//const KEY_H = 0x0b // Keyboard h and H
//const KEY_I = 0x0c // Keyboard i and I
//const KEY_J = 0x0d // Keyboard j and J
//const KEY_K = 0x0e // Keyboard k and K
//const KEY_L = 0x0f // Keyboard l and L
//const KEY_M = 0x10 // Keyboard m and M
//const KEY_N = 0x11 // Keyboard n and N
//const KEY_O = 0x12 // Keyboard o and O
//const KEY_P = 0x13 // Keyboard p and P
//const KEY_Q = 0x14 // Keyboard q and Q
//const KEY_R = 0x15 // Keyboard r and R
//const KEY_S = 0x16 // Keyboard s and S
//const KEY_T = 0x17 // Keyboard t and T
//const KEY_U = 0x18 // Keyboard u and U
//const KEY_V = 0x19 // Keyboard v and V
//const KEY_W = 0x1a // Keyboard w and W
//const KEY_X = 0x1b // Keyboard x and X
//const KEY_Y = 0x1c // Keyboard y and Y
//const KEY_Z = 0x1d // Keyboard z and Z

const KEY_1 = 0x1e // Keyboard 1 and !
const KEY_2 = 0x1f // Keyboard 2 and @
const KEY_3 = 0x20 // Keyboard 3 and #
const KEY_4 = 0x21 // Keyboard 4 and $
const KEY_5 = 0x22 // Keyboard 5 and %
const KEY_6 = 0x23 // Keyboard 6 and ^
const KEY_7 = 0x24 // Keyboard 7 and &
const KEY_8 = 0x25 // Keyboard 8 and *
const KEY_9 = 0x26 // Keyboard 9 and (
const KEY_0 = 0x27 // Keyboard 0 and )

const KEY_ENTER = 0x28      // Keyboard Return (ENTER)
const KEY_ESC = 0x29        // Keyboard ESCAPE
const KEY_BACKSPACE = 0x2a  // Keyboard DELETE (Backspace)
const KEY_TAB = 0x2b        // Keyboard Tab
const KEY_SPACE = 0x2c      // Keyboard Spacebar
const KEY_MINUS = 0x2d      // Keyboard - and _
const KEY_EQUAL = 0x2e      // Keyboard = and +
const KEY_LEFTBRACE = 0x2f  // Keyboard [ and {
const KEY_RIGHTBRACE = 0x30 // Keyboard ] and }
const KEY_BACKSLASH = 0x31  // Keyboard \ and |
const KEY_HASHTILDE = 0x32  // Keyboard Non-US # and ~
const KEY_SEMICOLON = 0x33  // Keyboard ; and :
const KEY_APOSTROPHE = 0x34 // Keyboard ' and "
const KEY_GRAVE = 0x35      // Keyboard ` and ~
const KEY_COMMA = 0x36      // Keyboard , and <
const KEY_DOT = 0x37        // Keyboard . and >
const KEY_SLASH = 0x38      // Keyboard / and ?
const KEY_CAPSLOCK = 0x39   // Keyboard Caps Lock

const KEY_F1 = 0x3a  // Keyboard F1
const KEY_F2 = 0x3b  // Keyboard F2
const KEY_F3 = 0x3c  // Keyboard F3
const KEY_F4 = 0x3d  // Keyboard F4
const KEY_F5 = 0x3e  // Keyboard F5
const KEY_F6 = 0x3f  // Keyboard F6
const KEY_F7 = 0x40  // Keyboard F7
const KEY_F8 = 0x41  // Keyboard F8
const KEY_F9 = 0x42  // Keyboard F9
const KEY_F10 = 0x43 // Keyboard F10
const KEY_F11 = 0x44 // Keyboard F11
const KEY_F12 = 0x45 // Keyboard F12

const KEY_SYSRQ = 0x46      // Keyboard Print Screen
const KEY_SCROLLLOCK = 0x47 // Keyboard Scroll Lock
const KEY_PAUSE = 0x48      // Keyboard Pause
const KEY_INSERT = 0x49     // Keyboard Insert
const KEY_HOME = 0x4a       // Keyboard Home
const KEY_PAGEUP = 0x4b     // Keyboard Page Up
const KEY_DELETE = 0x4c     // Keyboard Delete Forward
const KEY_END = 0x4d        // Keyboard End
const KEY_PAGEDOWN = 0x4e   // Keyboard Page Down
const KEY_RIGHT = 0x4f      // Keyboard Right Arrow
const KEY_LEFT = 0x50       // Keyboard Left Arrow
const KEY_DOWN = 0x51       // Keyboard Down Arrow
const KEY_UP = 0x52         // Keyboard Up Arrow

type KeyInstruction struct {
	modifier int
	padding  int
	key1     int
	key2     int
	key3     int
	key4     int
	key5     int
	key6     int
}

func print(num int) string {
	if num == 0 {
		return "\\0"
	}
	return fmt.Sprintf("\\x%x", num)
}

func (keyInstruction KeyInstruction) InstructionToString() string {
	return print(keyInstruction.modifier) +
		print(keyInstruction.padding) +
		print(keyInstruction.key1) +
		print(keyInstruction.key2) +
		print(keyInstruction.key3) +
		print(keyInstruction.key4) +
		print(keyInstruction.key5) +
		print(keyInstruction.key6)
}

func (keyInstruction *KeyInstruction) UnsetCtrl() *KeyInstruction {
	if keyInstruction.modifier&KEY_MOD_LCTRL == KEY_MOD_LCTRL {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_LCTRL
	}
	if keyInstruction.modifier&KEY_MOD_RCTRL == KEY_MOD_RCTRL {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_RCTRL
	}
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressCtrl() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_LCTRL
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressRightCtrl() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_RCTRL
	return keyInstruction
}

func (keyInstruction *KeyInstruction) UnsetAlt() *KeyInstruction {
	if keyInstruction.modifier&KEY_MOD_LALT == KEY_MOD_LALT {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_LALT
	}
	if keyInstruction.modifier&KEY_MOD_RALT == KEY_MOD_RALT {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_RALT
	}
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressAlt() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_LALT
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressRightAlt() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_RALT
	return keyInstruction
}

func (keyInstruction *KeyInstruction) UnsetShift() *KeyInstruction {
	if keyInstruction.modifier&KEY_MOD_LSHIFT == KEY_MOD_LSHIFT {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_LSHIFT
	}
	if keyInstruction.modifier&KEY_MOD_RSHIFT == KEY_MOD_RSHIFT {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_RSHIFT
	}
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressShift() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_LSHIFT
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressRightShift() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_RSHIFT
	return keyInstruction
}

func (keyInstruction *KeyInstruction) UnsetMeta() *KeyInstruction {
	if keyInstruction.modifier&KEY_MOD_LMETA == KEY_MOD_LMETA {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_LMETA
	}
	if keyInstruction.modifier&KEY_MOD_RMETA == KEY_MOD_RMETA {
		keyInstruction.modifier = keyInstruction.modifier - KEY_MOD_RMETA
	}
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressMeta() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_LMETA
	return keyInstruction
}

func (keyInstruction *KeyInstruction) PressRightMeta() *KeyInstruction {
	keyInstruction.modifier = keyInstruction.modifier | KEY_MOD_RMETA
	return keyInstruction
}

func (keyInstruction *KeyInstruction) SetKey1(key int) *KeyInstruction {
	keyInstruction.key1 = key
	return keyInstruction
}
func (keyInstruction *KeyInstruction) SetKey2(key int) *KeyInstruction {
	keyInstruction.key2 = key
	return keyInstruction
}
func (keyInstruction *KeyInstruction) SetKey3(key int) *KeyInstruction {
	keyInstruction.key3 = key
	return keyInstruction
}
func (keyInstruction *KeyInstruction) SetKey4(key int) *KeyInstruction {
	keyInstruction.key4 = key
	return keyInstruction
}
func (keyInstruction *KeyInstruction) SetKey5(key int) *KeyInstruction {
	keyInstruction.key5 = key
	return keyInstruction
}
func (keyInstruction *KeyInstruction) SetKey6(key int) *KeyInstruction {
	keyInstruction.key6 = key
	return keyInstruction
}

func Empty() *KeyInstruction {
	return &KeyInstruction{}
}
func (keyInstruction *KeyInstruction) Self() KeyInstruction {
	return *keyInstruction
}

func InstructionForBackspace() []KeyInstruction {
	var keyInstructions []KeyInstruction
	var key = (&KeyInstruction{}).
		SetKey1(KEY_BACKSPACE)
	keyInstructions = append(keyInstructions, *key)
	keyInstructions = append(keyInstructions, Empty().Self())
	return keyInstructions
}
func InstructionForEnter() []KeyInstruction {
	var keyInstructions []KeyInstruction
	var key = (&KeyInstruction{}).
		SetKey1(KEY_ENTER)
	keyInstructions = append(keyInstructions, *key)
	keyInstructions = append(keyInstructions, Empty().Self())
	return keyInstructions
}

func InstructionForCAD() []KeyInstruction {
	var keyInstructions []KeyInstruction
	var key = (&KeyInstruction{}).
		PressCtrl().
		PressAlt().
		SetKey1(KEY_DELETE)
	keyInstructions = append(keyInstructions, *key)
	keyInstructions = append(keyInstructions, Empty().Self())
	return keyInstructions
}

func InstructionForChar(msg string) KeyInstruction {
	if len(msg) == 1 {
		char := []rune(msg)[0]

		if unicode.IsLetter(char) {
			key := KeyInstruction{}
			if unicode.IsUpper(char) {
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = int(char - 61)
			} else {
				key.key1 = int(char - 93)
			}
			return key
		} else if unicode.IsDigit(char) {
			key := KeyInstruction{}
			switch char {
			case '1':
				key.key1 = KEY_1
			case '2':
				key.key1 = KEY_2
			case '3':
				key.key1 = KEY_3
			case '4':
				key.key1 = KEY_4
			case '5':
				key.key1 = KEY_5
			case '6':
				key.key1 = KEY_6
			case '7':
				key.key1 = KEY_7
			case '8':
				key.key1 = KEY_8
			case '9':
				key.key1 = KEY_9
			case '0':
				key.key1 = KEY_0
			}
			return key
		} else {
			key := KeyInstruction{}
			switch char {
			case '`':
				key.key1 = KEY_GRAVE
			case '~':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_GRAVE
			case '!':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_1
			case '@':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_2
			case '#':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_3
			case '$':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_4
			case '%':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_5
			case '^':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_6
			case '&':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_7
			case '*':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_8
			case '(':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_9
			case ')':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_0
			case '-':
				key.key1 = KEY_MINUS
			case '_':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_MINUS
			case '=':
				key.key1 = KEY_EQUAL
			case '+':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_EQUAL
			case '[':
				key.key1 = KEY_LEFTBRACE
			case '{':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_LEFTBRACE
			case ']':
				key.key1 = KEY_RIGHTBRACE
			case '}':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_RIGHTBRACE
			case '\\':
				key.key1 = KEY_BACKSLASH
			case '|':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_BACKSLASH
			case ';':
				key.key1 = KEY_SEMICOLON
			case ':':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_SEMICOLON
			case '\'':
				key.key1 = KEY_APOSTROPHE
			case '"':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_APOSTROPHE
			case ',':
				key.key1 = KEY_COMMA
			case '<':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_COMMA
			case '.':
				key.key1 = KEY_DOT
			case '>':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_DOT
			case '/':
				key.key1 = KEY_SLASH
			case '?':
				key.modifier = KEY_MOD_LSHIFT
				key.key1 = KEY_SLASH
			}
			return key
		}

	} else {
		panic("not a char: " + msg)
	}

}
func InstructionForString(msg string) []KeyInstruction {
	var keyInstructions []KeyInstruction
	if len(msg) == 0 {
		return keyInstructions
	}
	for _, char := range msg {
		keyInstructions = append(keyInstructions, InstructionForChar(string(char)))
		keyInstructions = append(keyInstructions, Empty().Self())
	}
	return keyInstructions
}

func typing(instruction KeyInstruction) {
	cmd := exec.Command("echo", "-ne", instruction.InstructionToString())
	outfile, err := os.Create("/dev/hidg0")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
}

func TryType(instructions []KeyInstruction, printMessage bool) {
	if instructions == nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	for _, x := range instructions {
		time.Sleep(time.Duration(15+rand.Intn(55)) * time.Millisecond)
		if printMessage {
			fmt.Println(x.InstructionToString())
		} else {
			typing(x)
		}
	}
}
func CastInstructionFromScript(msg string) *KeyInstruction {
	keyPairs := strings.Split(msg, "||")
	if len(keyPairs) > 2 || len(keyPairs) == 0 {
		panic("key paris format incorrect: " + msg)
	}
	key := Empty()
	if len(keyPairs[0]) < 2 {
		panic("function key format incorrect: " + keyPairs[0])
	}
	for _, funcKey := range strings.Split(keyPairs[0], ",") {
		funcKey = strings.Trim(funcKey, " ")
		switch {
		case funcKey == "ctrl":
			key.PressCtrl()
		case funcKey == "lctrl":
			key.PressCtrl()
		case funcKey == "rctrl":
			key.PressRightCtrl()
		case funcKey == "alt":
			key.PressAlt()
		case funcKey == "lalt":
			key.PressAlt()
		case funcKey == "ralt":
			key.PressRightAlt()
		case funcKey == "shift":
			key.PressShift()
		case funcKey == "lshift":
			key.PressShift()
		case funcKey == "rshift":
			key.PressRightShift()
		case funcKey == "meta":
			key.PressMeta()
		case funcKey == "lmeta":
			key.PressMeta()
		case funcKey == "rmeta":
			key.PressRightMeta()
		default:
			panic("function key expected: " + funcKey)
		}
	}
	if len(keyPairs) == 2 {
		pairs := strings.Split(keyPairs[1], ",")
		var valid = true
		for i, str := range pairs {
			pairs[i] = strings.Trim(str, " ")
			if len(pairs[i]) != 3 || pairs[i][0] != '{' && pairs[i][2] != '}' {
				valid = false
				break
			}
			pairs[i] = string(pairs[i][1])
		}
		if len(pairs) > 6 {
			valid = false
		}
		if !valid {
			panic("invalid keys: " + keyPairs[1])
		}
		for i, str := range pairs {
			k := InstructionForChar(str)
			switch i {
			case 0:
				key.SetKey1(k.key1)
			case 1:
				key.SetKey2(k.key1)
			case 2:
				key.SetKey3(k.key1)
			case 3:
				key.SetKey4(k.key1)
			case 4:
				key.SetKey5(k.key1)
			case 5:
				key.SetKey6(k.key1)
			}
		}
	}
	return key
}
func CastFromScript(msg string) []KeyInstruction {
	keys := []KeyInstruction{}
	switch msg[0] {
	case '|':
		keys = append(keys, *CastInstructionFromScript(msg[1:]))
		keys = append(keys, Empty().Self())
	case ':':
		for _, key := range InstructionForString(msg[1:]) {
			keys = append(keys, key)
		}
	case '<':
		keys = append(keys, *CastInstructionFromScript(msg[1:]))
	case '-':
		keys = append(keys, Empty().Self())
	default:
		return nil
	}
	return keys
}

func TryTypeFromBuffer(reader bufio.Reader, printMessage bool) {
	rand.Seed(time.Now().UnixNano())

	for {
		bytes, _, err := reader.ReadLine()
		time.Sleep(time.Duration(15+rand.Intn(55)) * time.Millisecond)
		msg := string(bytes)

		if err == io.EOF {
			return
		} else if err != nil {
			panic(err)
		} else if len(msg) < 2 {
			continue
		} else {
			TryType(CastFromScript(msg), printMessage)
		}

	}
}
