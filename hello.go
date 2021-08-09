package main

import "fmt"
import "github.com/pulysak/go_hello_world/morestrings"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
}
