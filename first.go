package main

import {
	"fmt"
	"github.com/irlac/first-go/morestrings"
	"github.com/google/go-cmp/cmp"
}

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}