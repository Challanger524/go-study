package main

import (
	"fmt"
)

// Strings, bytes, runes and characters in Go:
// https://go.dev/blog/strings
func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)
	fmt.Printf("% x\n", sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Println()
	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)

	fmt.Println("\nRange for loop reveals the `runes`:")
	const nihongo = "日本語"
	for idx, runeval := range nihongo {
		fmt.Printf("Rune at byte pos [%02d]: %#U\n", idx, runeval)
	}
}
