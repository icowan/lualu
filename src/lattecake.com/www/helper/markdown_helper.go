package helper

import (
	"strings"
	"log"
)

type GoMarkdown struct {
	DefinitionData []string
	Lines          []string
}

const version = "1.0.1"

var (
	charlist = []string{
		"\0",
		"\t",
		"\n",
		"\r",
		" ",
	}
)

func (m *GoMarkdown) Text(text string) string {
	text = strings.Replace(text, "\r\n", "\n", -1)
	text = strings.Replace(text, "\r", "\n", -1)

	text = strings.Trim(text, "\n")

	m.Lines = strings.Split(text, "\n")

	markup := m.RunLines()

	return strings.Trim(markup, "\n")
}

func (m *GoMarkdown) RunLines() string {

	var currentBlock map[string]interface{}
	for _, line := range m.Lines {
		if chop(line) == "" {
			if len(currentBlock) > 0 {
				currentBlock["nterrupted"] = true
			}
			continue
		}

		if UnicodeIndex(line, "\t") != -1 {
			parts := strings.Split(line, "\t")
			line = parts[0]

			for _, part := range parts[1:] {
				shortage := 4 - len([]rune(string(line)))%4
				line += strings.Repeat(" ", shortage)
				line += part
			}
		}

		indent := 0

		log.Println(indent)
	}

	return string("")
}

func chop(text string) string {
	for _, item := range charlist {
		text = strings.Replace(text, item, "", -1)
	}

	return text
}

func UnicodeIndex(text string, substr string) int {
	res := strings.Index(text, substr)

	if res != -1 {
		prefix := []byte(text)[0:res]
		rs := []rune(string(prefix))
		res = len(rs)
	}

	return res
}
