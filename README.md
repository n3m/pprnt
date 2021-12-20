# pprnt [![travis](https://travis-ci.com/n3m/pprnt.svg?branch=master)](https://travis-ci.com/github/n3m/pprnt) [![gopls](https://pkg.go.dev/badge/github.com/n3m/pprnt)](https://pkg.go.dev/github.com/n3m/pprnt) [![goreport](https://goreportcard.com/badge/github.com/n3m/pprnt)](https://goreportcard.com/report/github.com/n3m/pprnt)

Pretty Print for Golang (Maps, Models and Normal Variables)

## How to install

##### Version: 1.11.0

`go get github.com/n3m/pprnt`

## Description

Pretty Print is a Golang Package designed to print Maps and Structs as if they were JSON structures instead of the ugly print that default/external printing packages offer

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

Tired of printing structs or maps in go without any identation? Well, I was. I was tired of having to search for the key that I was looking for in the whole mess of the print.
Imagine that you have a really big golang struct and you want to compare data between the same variable after some changes.
Without identation it would take a lot of effort to find the keys that you're looking for, and I was tired of that.

TL;DR: I hate printing structs or maps in go, and I created a solution.
