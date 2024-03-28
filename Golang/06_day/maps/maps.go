package main

import "fmt"

func main() {
	fmt.Println("maps is all about key and pair")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"
	languages["r"] = "rust"

	fmt.Println("list of all languages: ", languages)

	fmt.Println("accessing values based on keywords: ", languages["r"])

	delete(languages, "rb")

	fmt.Println("list of all languages: ", languages)
}
