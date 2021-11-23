package main

import (
	"fmt"

	"github.com/skeptycal/gomake"
)

func main() {
	args := gomake.NewCLIArgs()
	fmt.Println(args)
}
