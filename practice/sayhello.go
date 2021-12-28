package main

func Hello(language string) string{
	if language == "Spanish"{
		return "Hola!"
	}
	switch language {
	case "Spanish":
		return "Hola!"
	case "French":
		return "Bonjour!"
	case "Sanskrit":
		return "Namaskaram"
	case "Tamil":
		return "Vanakkam"
	default:
		return "Hello!"
	}
}