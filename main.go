package main

import (
	"fmt"

	_ "github.com/zinclabs/zinc/pkg/handlers/auth"
	_ "github.com/zinclabs/zinc/pkg/handlers/document"
	_ "github.com/zinclabs/zinc/pkg/handlers/index"
	_ "github.com/zinclabs/zinc/pkg/handlers/search"
	_ "github.com/zinclabs/zinc/pkg/meta/elastic"
)

func main() {
	fmt.Println("hello world")
}
