package hello

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjour "

func Hello(name string, language string) string {

	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {

	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
