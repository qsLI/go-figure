package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"unicode"
	"log"
	"strconv"
	"io"
)

func main() {
	æs := ""
	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00E6'} {
		fmt.Printf("[0x%X '%c'] ", char, char)
		æs += string(char)
	}

	/*
	0  U+0076 'v' 76
	1  U+00E5 'å' C3 A5
	3  U+0074 't' 74
	4  U+0074 't' 74
	5  U+0020 ' ' 20
	6  U+006F 'o' 6F
	7  U+0067 'g' 67
	8  U+0020 ' ' 20
	9  U+0074 't' 74
	10 U+00F8 'ø' C3 B8
	12 U+0072 'r' 72
	13 U+0074 't' 74
	 */
	phrase := "vått og tørt"
	fmt.Printf("string: \"%s\"\n", phrase)
	fmt.Println("index rune char bytes")
	for index, char := range phrase {
		fmt.Printf("%-2d %U '%c' % X\n",
			index, char, char,
			[]byte(string(char)))
	}

	/**
	go do not support index style, do not use this way!
	0: v
	1: Ã
	2: ¥
	3: t
	4: t
	5:
	6: o
	7: g
	8:
	9: t
	10: Ã
	11: ¸
	12: r
	13: t
	 */
	for i := 0; i < len(phrase); i++ {
		fmt.Printf("%d: %c\n", i,phrase[i])
	}

	r, size := utf8.DecodeRuneInString(phrase)
	fmt.Println(r, size)

	/**
	convert to rune and using slice
	 */
	fmt.Println("using rune..")
	rr := []rune(phrase)
	for i := 0; i < len(rr); i++ {
		fmt.Printf("%d: %c\n", i, rr[i])
	}

	asciiIndex()

	unicodeIndex() // Prints: rå vær

	fmt.Println("\nstring reader")
	reader := strings.NewReader("Café")
	for {
		rrr, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else {
			log.Println(err)
		}
		fmt.Printf("size=%d, %s\n", size, string(rrr))
	}

	// strconv
	b, e := strconv.ParseBool("true")
	if e == nil {
		fmt.Printf("%t %v\n", b, b)
	}

	grades := []float64{99, 98, 96}
	for _, grade := range grades {
		grade *= 1.5
	}
	fmt.Println(grades) // unchanged

	for i := range grades {
		grades[i] *= 1.5
	}
	fmt.Println(grades)


}
func unicodeIndex() {
	fmt.Printf("\n\nusing unicode index\n")
	line := "rå tørt\u2028vær"
	i := strings.IndexFunc(line, unicode.IsSpace)
	// i == 3
	firstWord := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace)
	// j == 9
	r, size := utf8.DecodeRuneInString(line[j:])
	fmt.Printf("% X\n", []byte(string(r)))
	// size == 3
	lastWord := line[j+size:]
	// j + size == 12
	fmt.Println(firstWord, lastWord)
}

func asciiIndex() {
	fmt.Printf("\n\nusing ascii index\n")
	line := "røde og gule sløjfer"
	i := strings.Index(line, " ")
	// Get the index of the first space
	firstWord := line[:i]
	// Slice up to the first space
	j := strings.LastIndex(line, " ")
	// Get the index of the last space
	lastWord := line[j+1:]
	// Slice from after the last space
	fmt.Println(firstWord, lastWord)
	// Prints: røde sløjfe
}
