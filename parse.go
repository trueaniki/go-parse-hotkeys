package parsehotkeys

import (
	"fmt"
	"strings"

	"golang.design/x/hotkey"
)

func ParseKey(key string) (hotkey.Key, error) {
	key = strings.ToLower(key)
	if k, ok := KeysMap[key]; ok {
		return k, nil
	}
	return hotkey.Key(0), fmt.Errorf("unknown key: %s", key)
}

func ParseMod(mod string) (hotkey.Modifier, error) {
	mod = strings.ToLower(mod)
	if m, ok := ModsMap[mod]; ok {
		return m, nil
	}
	return 0, fmt.Errorf("unknown modifier: %s", mod)
}

func Parse(hk string, delimiter string) (*hotkey.Hotkey, error) {
	tokens := strings.Split(hk, delimiter)
	key, err := ParseKey(tokens[len(tokens)-1])
	if err != nil {
		return nil, err
	}
	modifiers := make([]hotkey.Modifier, 0, len(tokens)-1)
	for _, token := range tokens[:len(tokens)-1] {
		modifier, err := ParseMod(token)
		if err != nil {
			return nil, err
		}
		modifiers = append(modifiers, modifier)
	}
	return hotkey.New(modifiers, key), nil
}
