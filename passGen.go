package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const lowerLetterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
const upperLetterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const symbolBytes = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {

	digitFlag := flag.Int("d", 8, "Number of password digits.")
	symbolFlag := flag.Bool("s", false, "Include symbols.")
	caseFlag := flag.Bool("u", false, "Include upper case digits.")
	flag.Parse()

	fmt.Printf("\nDigits: %d\nSymbols: %t\nUpper case: %t\n", *digitFlag, *symbolFlag, *caseFlag)
	var charset string
	if *symbolFlag == true && *caseFlag == false {
		charset = lowerLetterBytes + symbolBytes
	} else if *symbolFlag == false && *caseFlag == true {
		charset = lowerLetterBytes + upperLetterBytes

	} else if *symbolFlag == true && *caseFlag == true {
		charset = lowerLetterBytes + upperLetterBytes + symbolBytes

	} else {
		charset = lowerLetterBytes

	}
	output := StringWithCharset(*digitFlag, charset)
	fmt.Printf("\nPassword: %s\n\n", output)

}
