package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kr/text"
	"github.com/mattes/go-asciibot"
)

func main() {
	var maxWidth int
	var botSeed string
	flag.IntVar(&maxWidth, "width", 40, "Approximate line width (for wrapping).")
	flag.StringVar(&botSeed, "bot", "", "5-character bot seed.")
	flag.Parse()

	PrintBubble(FlattenIntoLines(WrapParagraphs(maxWidth), maxWidth))
	PrintBot(GenerateBot(botSeed))
}

func WrapParagraphs(lineLength int) (paragraphs []string) {
	all, _ := ioutil.ReadAll(os.Stdin)
	all = bytes.ReplaceAll(all, []byte{'\t'}, []byte("    "))
	paragraphs = strings.Split(string(all), "\n\n")
	for x := range paragraphs {
		paragraphs[x] = text.Wrap(paragraphs[x], lineLength)
	}
	return paragraphs
}

func FlattenIntoLines(paragraphs []string, maxWidth int) (lines []string, lineLength int) {
	if len(paragraphs) == 1 && len(paragraphs[0]) < maxWidth {
		return append(lines, paragraphs[0]), len(paragraphs[0])
	}
	for p, paragraph := range paragraphs {
		lines = append(lines, strings.Split(paragraph, "\n")...)
		if p < len(paragraphs)-1 {
			lines = append(lines, strings.Repeat(" ", lineLength))
		}
	}
	max := 9
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}

	// fill in all lines with trailing spaces
	for l := range lines {
		lines[l] += strings.Repeat(" ", max-len(lines[l]))
	}

	// get rid of trailing blank lines
	for strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-2]
	}
	return lines, max
}

func PrintBubble(lines []string, lineLength int) {
	fmt.Println(" " + strings.Repeat("_", lineLength+2))
	for l, line := range lines {
		if len(lines) == 1 {
			fmt.Println("< " + line + " >")
		} else if l == 0 {
			fmt.Println("/ " + line + " \\")
		} else if l == len(lines)-1 {
			fmt.Println("\\ " + line + " /")
		} else {
			fmt.Println("| " + line + " |")
		}
	}
	fmt.Println(" " + strings.Repeat("-", lineLength+2))
}

func GenerateBot(botSeed string) string {
	if bot, err := asciibot.Generate(botSeed); err == nil {
		return bot
	} else {
		return asciibot.Random()
	}
}

func PrintBot(bot string) {
	for x, line := range strings.Split(bot, "\n") {
		if x == 0 {
			fmt.Println(`        \ ` + line)
		} else if x == 1 {
			fmt.Println(`         \` + line)
		} else {
			fmt.Println("          " + line)
		}
	}
}
