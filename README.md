# BBAN generator
> Simple BBAN generator & validator

[![Go Report Card](https://goreportcard.com/badge/github.com/m1ome/bban_gen)](https://goreportcard.com/report/github.com/m1ome/bban_gen)
[![GoDoc](https://godoc.org/github.com/m1ome/bban_gen?status.svg)](https://godoc.org/github.com/m1ome/bban_gen)
[![Build Status](https://travis-ci.org/m1ome/bban_gen.svg?branch=master)](https://travis-ci.org/m1ome/bban_gen)
[![Coverage Status](https://coveralls.io/repos/github/m1ome/bban_gen/badge.svg?branch=master)](https://coveralls.io/github/m1ome/bban_gen?branch=master)

## Installation
```go
go get github.com/m1ome/bban_gen
```

## Usage
```go
package main

import (
    "fmt"

    bban "github.com/m1ome/bban_gen"
)

func main() {
	account := bban.Random("040577", "13439317554524", bban.DoubleMod)
	fmt.Printf("Generated account: %s\n", account)

	next := bban.Next("040577", account, "13439317554524", bban.DoubleMod)
	fmt.Printf("Next account: %s\n", next)

	fmt.Printf("Validity passing: %v\n", bban.Validate("040577", next, "13439317554524", bban.DoubleMod))
}
```