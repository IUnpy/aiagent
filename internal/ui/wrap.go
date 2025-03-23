package ui

import "strings"

// ChineseWrap 处理中文文本换行
func ChineseWrap(text string, maxCharsPerLine int) string {
	runes := []rune(text)
	var result strings.Builder

	for i, r := range runes {
		result.WriteRune(r)
		if (i+1)%maxCharsPerLine == 0 && i != len(runes)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}
