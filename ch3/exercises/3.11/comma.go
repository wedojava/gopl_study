package comma3

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if len(s) > 1 && s[0] == '-' {
		buf.WriteRune('-')
		s = s[1:]
	}

	i := strings.Index(s, ".")
	if i != -1 {
		buf.WriteString(commaInt(s[:i]))
		buf.WriteRune('.')
		buf.WriteString(s[i+1:])
	} else {
		buf.WriteString(commaInt(s))
	}

	return buf.String()
}

func commaInt(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i, c := range s {
		buf.WriteRune(c)
		if ((1+i-n)%3) == 0 && (i != n-1) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
