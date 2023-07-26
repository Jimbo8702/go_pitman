package main

import "strings"

func trimHTML(html string) string {
	html = strings.ReplaceAll(html, " ", "")
	// Remove newlines
	html = strings.ReplaceAll(html, "\n", "")
	
	return html
}