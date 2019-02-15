package as

import (
	"fmt"
	"strings"
)

func ParseBool(inp string) (value, ok bool) {
	ok = true

	switch strings.ToLower(strings.TrimSpace(inp)) {
	case "true", "yes", "on", "1", "enable", "enabled":
		value = true
	case "false", "no", "off", "0", "disable", "disabled":
		value = false
	default:
		ok = false
	}
	return
}

func IntToBool(inp int) (value, ok bool) {
	ok = (inp == 0) || (inp == 1)
	if ok {
		value = inp == 1
	}
	return
}

func FloatToBool(inp float64) (value, ok bool) {
	ok = (inp == 0) || (inp == 1)
	if ok {
		value = inp == 1
	}
	return
}

func Bool(inp interface{}) (value, ok bool) {
	ok = true

	switch v := inp.(type) {
	case bool:
		value = v
	case string:
		value, ok = ParseBool(v)
	case fmt.Stringer:
		value, ok = ParseBool(v.String())
	case int:
		value, ok = IntToBool(int(v))
	case uint:
		value, ok = IntToBool(int(v))
	case int8:
		value, ok = IntToBool(int(v))
	case uint8:
		value, ok = IntToBool(int(v))
	case int16:
		value, ok = IntToBool(int(v))
	case uint16:
		value, ok = IntToBool(int(v))
	case int32:
		value, ok = IntToBool(int(v))
	case uint32:
		value, ok = IntToBool(int(v))
	case int64:
		value, ok = IntToBool(int(v))
	case uint64:
		value, ok = IntToBool(int(v))
	case uintptr:
		value, ok = IntToBool(int(v))
	case float32:
		value, ok = FloatToBool(float64(v))
	case float64:
		value, ok = FloatToBool(v)
	default:
		ok = false
	}

	return
}
