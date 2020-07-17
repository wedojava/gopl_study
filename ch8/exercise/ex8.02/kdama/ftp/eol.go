package ftp

// In ASCII mode, the end-of-line sequence should be CRLF.
// https://tools.ietf.org/html/rfc959#page-19
// eol is end-of-line sequence
func (c *Conn) eol() string {
	switch c.datatype {
	case ascii:
		return "\r\n"
	case image:
		return "\n"
	default:
		return "\n"
	}
}
