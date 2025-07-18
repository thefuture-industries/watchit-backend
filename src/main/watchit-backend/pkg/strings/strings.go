package strings

import (
	"fmt"
	"strconv"
	"strings"
)

func Formatted(b *strings.Builder, format string, args ...any) {
	argIndex := 0

	for i := 0; i < len(format); i++ {
		if format[i] == '%' && i+1 < len(format) {
			i++
			if argIndex >= len(args) {
				continue
			}

			switch format[i] {
			case 'd':
				b.WriteString(strconv.FormatInt(ToInt64(args[argIndex]), 10))
			case 'x':
				if bs, ok := args[argIndex].([]byte); ok {
					b.WriteString(strings.ToLower(HexEncode(bs)))
				} else {
					b.WriteString("<?>")
				}
			case 's':
				b.WriteString(args[argIndex].(string))
			case 'v':
				b.WriteString(fmt.Sprint(args[argIndex]))
			default:
				b.WriteByte('%')
				b.WriteByte(format[i])
			}

			argIndex++
		} else {
			b.WriteByte(format[i])
		}
	}
}

func ToInt64(v any) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	default:
		return 0
	}
}

func HexEncode(data []byte) string {
	const hex = "0123456789abcdef"

	var b strings.Builder

	for _, c := range data {
		b.WriteByte(hex[c>>4])
		b.WriteByte(hex[c&0x0f])
	}

	return b.String()
}
