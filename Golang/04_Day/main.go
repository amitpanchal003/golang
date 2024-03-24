package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to day 4 of golang series")
	fmt.Println("please rate our pizza betwwen 1 to 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("thans for your feedback,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strconv is package which is used for conversion.
	//parseFloat is method of that package. it require two parameter which is targetValue and base of that value.
	//strings is also a package which provide lot of methods.
	//TrimSpace is one of them.
	// this will remove the leading space of string
	// String and strings different.

	// error handling
	if err != nil { // we are saying that if there is something.
		fmt.Println(err) // just print that(err)
	} else { // otherwise
		fmt.Println("Added 1 to your rathing,", numRating+1)
		// perform the given operation and just print that
	}

}
