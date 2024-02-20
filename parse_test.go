package parsehotkeys_test

import (
	"fmt"
	"testing"

	parsehotkeys "github.com/trueaniki/go-parse-hotkeys"
)

func TestParse(t *testing.T) {
	input := "ctrl+shift+a"
	hk, err := parsehotkeys.Parse(input, "+")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Hotkey: %s\n", hk)

	input = "CTRL alt a"
	hk, err = parsehotkeys.Parse(input, " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Hotkey: %s\n", hk)

	input = "Ctrl&Shift&A"
	hk, err = parsehotkeys.Parse(input, "&")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Hotkey: %s\n", hk)
}
