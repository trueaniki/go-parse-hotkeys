package parsehotkeys

import (
	"fmt"
	"strings"

	"golang.design/x/hotkey"
)

func parseKey(key string) (hotkey.Key, error) {
	if k, ok := KeysMap[key]; ok {
		return k, nil
	}
	return hotkey.Key(0), fmt.Errorf("unknown key: %s", key)
}

func parseModifier(modifier string) (hotkey.Modifier, error) {
	if m, ok := ModsMap[modifier]; ok {
		return m, nil
	}
	return 0, fmt.Errorf("unknown modifier: %s", modifier)
}

func Parse(hk string, delimiter string) (*hotkey.Hotkey, error) {
	tokens := strings.Split(hk, delimiter)
	key, err := parseKey(tokens[len(tokens)-1])
	if err != nil {
		return nil, err
	}
	modifiers := make([]hotkey.Modifier, 0, len(tokens)-1)
	for _, token := range tokens[:len(tokens)-1] {
		modifier, err := parseModifier(token)
		if err != nil {
			return nil, err
		}
		modifiers = append(modifiers, modifier)
	}
	return hotkey.New(modifiers, key), nil
}
