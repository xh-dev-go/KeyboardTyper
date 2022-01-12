package KeyboardTyper

import "testing"

func TestClear(t *testing.T) {
	if Empty().InstructionToString() != "\\0\\0\\0\\0\\0\\0\\0\\0" {
		t.Error("clear  not matched")
	}
}

func TestCAD(t *testing.T) {
	if Empty().PressCtrl().PressAlt().SetKey1(KEY_DELETE).InstructionToString() != "\\x5\\0\\x4c\\0\\0\\0\\0\\0" {
		t.Error("clear  not matched")
	}
}

func TestCastInstructionFromScript(t *testing.T) {
	var key = CastInstructionFromScript("ctrl")
	if key.InstructionToString() != "\\x1\\0\\0\\0\\0\\0\\0\\0" {
		t.Error("error")
	}

	key = CastInstructionFromScript("ctrl, lalt")
	if key.InstructionToString() != "\\x5\\0\\0\\0\\0\\0\\0\\0" {
		t.Error("error")
	}

	key = CastInstructionFromScript("ctrl, lalt ||{A}")
	if key.InstructionToString() != "\\x5\\0\\x4\\0\\0\\0\\0\\0" {
		t.Error("error")
	}

	key = CastInstructionFromScript("ctrl, lalt ||{A},{a},{B},{b},{C} ,{c}")
	if key.InstructionToString() != "\\x5\\0\\x4\\x4\\x5\\x5\\x6\\x6" {
		t.Error("error")
	}
}
