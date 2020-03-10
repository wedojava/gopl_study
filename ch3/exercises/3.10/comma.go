package comma2

import "bytes"

func comma(s string) string {
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
