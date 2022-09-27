# go-swiss-knife

[![Go Reference](https://pkg.go.dev/badge/github.com/matribots/go-swiss-knife.svg)](https://pkg.go.dev/github.com/matribots/go-swiss-knife)

Just like a swiss knife - A simple go utility toolkits which is used for general purpose during daily coding.

## Char&String Knives

Locate in `charstring` package. This package is mainly used for dealing with strings and chars misc stuffs.

Here is the example:

```go
package main

import "github.com/matribots/go-swiss-knife/charstring"

// you must defined ascii type first
type ascii byte

func main() {
    var len = 20
    asciiChars := make([]ascii, len)

    // Generate 20 random printable ascii chars that locate in [28,125)
    charstring.RandomPrintableASCII(28, 125, len)


    // Check if a string is empty
    str1, str2, str3, str4 :="", "    ", "NOT_EMPTY", "NOT EMPTY"
    // str1 is empty
    charstring.IsEmptyString(str1)
    // str2 is empty
    charstring.IsEmptyString(str2)
    // str3 is not empty
    charstring.IsEmptyString(str3)
    // str4 is not empty
    charstring.IsEmptyString(str4)
}
```

## Crypto Knives

Locate in `crypto` package. This package is mainly used for dealing password, encryption and decryption miscs.

Here is the example:

```go
package main

import "github.com/matribots/go-swiss-knife/crypto"

func main(){
    // Generate a random strong password that length is 20.
    crypto.RandomPassword(20)    
}
```

## Numbers Knives

Locate in `numbers` package. This package is mainly used for dealing numbers related miscs.

Here is the example:

```go
package main

import "github.com/matribots/go-swiss-knife/numbers"

func main(){
    // Generate a random postive int that locate in [0,80).
    numbers.RandomPostiveInt(80)
    // Generate 20 random postive ints that locate in [0,100).
    numbers.RandomPostiveInts(0, 100, 20)
}
```
