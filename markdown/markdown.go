package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	header := 0
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	var pos int
	var listLength int
	var listOpened bool
	var html string
	var displayHash bool
	for pos < len(markdown) {
		char := markdown[pos]
		switch {
		case char == '#':
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			if header == 7 {
				html += fmt.Sprintf("<p>%s ", strings.Repeat("#", header))
			} else if displayHash {
				html += "# "
				header--
			} else {
				html += fmt.Sprintf("<h%d>", header)
			}
		case char == '*' && header == 0 && strings.Contains(markdown, "\n"):
			if listLength == 0 {
				html += "<ul>"
			}
			listLength++
			if !listOpened {
				html += "<li>"
				listOpened = true
			} else {
				html += string(char) + " "
			}
			pos++
		case char == '\n':
			if listOpened && strings.LastIndex(markdown, "\n") == pos && strings.LastIndex(markdown, "\n") > strings.LastIndex(markdown, "*") {
				html += "</li></ul><p>"
				listOpened = false
				listLength = 0
			} else if listLength > 0 && listOpened {
				html += "</li>"
				listOpened = false
			} else if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		default:
			html += string(char)
		}
		displayHash = true
		pos++
	}
	switch {
	case header == 7:
		return fmt.Sprintf("%s</p>", html)
	case header > 0:
		return fmt.Sprintf("%s</h%d>", html, header)
	case listLength > 0:
		return fmt.Sprintf("%s</li></ul>", html)
	case strings.Contains(html, "<p>"):
		return fmt.Sprintf("%s</p>", html)
	default:
		return fmt.Sprintf("<p>%s</p>", html)
	}
}
