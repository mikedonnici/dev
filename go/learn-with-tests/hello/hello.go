package hello

import (
	"fmt"
	"strings"
)

const englishHello = "Hello, %s!"
const spanishHello = "Hola, %s!"
const frenchHello = "Bonjour, %s!"
const italianHello = "Bonjourno, %s!"

func Hello(s string, lang string) string {
	if s == "" {
		s = "World"
	}
	switch strings.ToLower(lang) {
	case "es", "spanish":
		return fmt.Sprintf(spanishHello, s)
	case "fr", "french":
		return fmt.Sprintf(frenchHello, s)
	case "it", "italian":
		return fmt.Sprintf(italianHello, s)
	}
	return fmt.Sprintf(englishHello, s)
}
