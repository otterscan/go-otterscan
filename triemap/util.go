package triemap

import (
	"encoding/hex"
	"strings"
)

func RawFromHex(s string) string {
	ans, _ := hex.DecodeString(s)
	return string(ans)
}

func SanityCheck(s string) string {
	if s == "" {
		return ""
	}
	if strings.Contains(s, "(") && strings.Contains(s, ")") {
		return s
	}
	return ""
}
