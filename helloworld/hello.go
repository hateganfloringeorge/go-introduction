package main

const (
	french    = "French"
	spanish   = "Spanish"
	taiwanese = "Taiwanese"

	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchHelloPrefix    = "Bonjour, "
	taiwaneseHelloPrefix = "Li ho, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case taiwanese:
		prefix = taiwaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}