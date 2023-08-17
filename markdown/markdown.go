package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	var pos int
	var listLength int
	var listOpened bool
	var html string
	var displayHash bool
	var headerTracker int
	for pos < len(markdown) {
		char := markdown[pos]
		switch {
		case char == '#':
			for char == '#' {
				headerTracker++
				pos++
				char = markdown[pos]
			}
			if headerTracker == 7 {
				html += fmt.Sprintf("<p>%s ", strings.Repeat("#", headerTracker))
			} else if displayHash {
				html += "# "
				headerTracker--
			} else {
				html += fmt.Sprintf("<h%d>", headerTracker)
			}
		case char == '*' && headerTracker == 0 && strings.Contains(markdown, "\n"):
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
			} else if headerTracker > 0 {
				html += fmt.Sprintf("</h%d>", headerTracker)
				headerTracker = 0
			}
		default:
			html += string(char)
		}
		displayHash = true
		pos++
	}
	switch {
	case headerTracker == 7:
		return fmt.Sprintf("%s</p>", html)
	case headerTracker > 0:
		return fmt.Sprintf("%s</h%d>", html, headerTracker)
	case listLength > 0:
		return fmt.Sprintf("%s</li></ul>", html)
	case strings.Contains(html, "<p>"):
		return fmt.Sprintf("%s</p>", html)
	default:
		return fmt.Sprintf("<p>%s</p>", html)
	}
}
