# pprnt [![travis](https://travis-ci.com/n3m/pprnt.svg?branch=master)](https://travis-ci.com/github/n3m/pprnt) [![gopls](https://pkg.go.dev/badge/github.com/n3m/pprnt)](https://pkg.go.dev/github.com/n3m/pprnt) [![goreport](https://goreportcard.com/badge/github.com/n3m/pprnt)](https://goreportcard.com/report/github.com/n3m/pprnt)

Pretty Print for Golang (Print anything with format basically)

## Get Started

##### Version: 1.12.0

`go get github.com/n3m/pprnt`

## Description

Designed to Pretty Print Stuff on your Terminal, pretty basic and synchronous. 

### Usage

- Example Code:

```go
package main

import "github.com/n3m/pprnt"

func main(){
  example := map[string]interface{}{"hell": map[string]interface{}{"yeah":"brother"}}

  // Print(interface) :: The function Print() expects an interface to be passed as parameters
  // So you can pass any type of variable
  pprint.Print(example)
}
```

- Example Output

```
$ go run .

{
   "hell": {
      "yeah": "brother"
   }
}
```

### Project Reason

TL;DR: I do not like the way that current printing packages handle Terminal Printing in Golang, so I decided to make my own.

Tired of printing structs or maps in go without any identation? Tired of using json.Marshal -> []bytes -> String to print with identation?
Well, I was. I was tired of having to search for the key that I was looking for in the whole mess of the print.
Imagine that you have a really big struct/map and you want to visually understand whats happening to the data; this can be a really time consuming task using something like fmt.Printf or log.Printf since those packages just print directly to the terminal; so in order to erradicate this awful printing style, I decided to create a package that could print in a "NodeJS"-style out of the box. That's the reason, believe it or not. I wasted 1 day of my life with this + some hours fixing bugs and testing it.

### Disclaimer

This package is designed to be ran on development environments. Beware of turning on using Pretty Printing from this package while on Production Mode, since it can generate performance issues, due to the use of recursion.
This package is released "as-is" and I'm not responsible for any liabilty that could result due to the use of this package.
