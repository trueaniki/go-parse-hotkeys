package parsehotkeys_test

import (
	"testing"

	parsehotkeys "github.com/trueaniki/go-parse-hotkeys"
	"golang.design/x/hotkey"
)

func TestParse(t *testing.T) {
	input := "ctrl+shift+a"
	hk, err := parsehotkeys.Parse(input, "+")
	if err != nil {
		t.Fatal(err)
	}
	if hk.String() != hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyA).String() {
		t.Error("Test failed")
	}

	input = "CTRL alt a"
	hk, err = parsehotkeys.Parse(input, " ")
	if err != nil {
		t.Fatal(err)
	}
	if hk.String() != hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModOption}, hotkey.KeyA).String() {
		t.Error("Test failed")
	}

	input = "Ctrl&Shift&A"
	hk, err = parsehotkeys.Parse(input, "&")
	if err != nil {
		t.Fatal(err)
	}
	if hk.String() != hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyA).String() {
		t.Error("Test failed")
	}
}
