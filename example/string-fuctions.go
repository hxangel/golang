//The standard library’s strings package provides many useful string-related functions. Here are some examples to give you a sense of the package.
package main
import (
	"fmt"
	"strings"
)
//We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println
func main() {
	//Here’s a sample of the functions available in strings. Note that these are all functions from package, not methods on the string object itself. This means that we need pass the string in question as the first argument to the function.
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	p()
	//You can find more functions in the strings package docs.
	//Not part of strings but worth mentioning here are the mechanisms for getting the length of a string and getting a character by index.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
/*
$ go run string-functions.go
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
toLower:    test
ToUpper:    TEST
Len:  5
Char: 101
*/
//Next example: String Formatting.