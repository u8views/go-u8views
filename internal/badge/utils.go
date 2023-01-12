package badge

import (
	"regexp"
	"strings"
)

func stripXmlWhitespace(xml string) string {
	return strings.TrimSpace(regexp.MustCompile(`<\s+`).ReplaceAllString(regexp.MustCompile(`>\s+`).ReplaceAllString(xml, ">"), "<"))
}
