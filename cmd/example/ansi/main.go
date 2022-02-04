package main

import (
	"fmt"

	"github.com/skeptycal/gomake"
)

func main() {
	a := gomake.AnsiString{}
	a.Set(15, 0, 0)

	fmt.Printf("This is an %sexample.", a)
}
