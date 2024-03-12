# Go Parse Hotkeys
A library to parse hotkeys from a string with given delimeter and return hotkeys in format of [golang.design/x/hotkey](https://pkg.go.dev/golang.design/x/hotkey)
## Usage
```go
hotkey := parsehotkeys.Parse("Ctrl+Shift+A", "+")
hotkey = parsehotkeys.Parse("CTRL OPTION A", " ")
hotkey = parsehotkeys.Parse("ctrl&alt&a", "&")
```
Then you can use according to [golang.design/x/hotkey](https://pkg.go.dev/golang.design/x/hotkey) docs
```go
hotkey.Register()
```

Full list of keys string representation is available at `map.go`
