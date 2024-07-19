package main

import (
	"github.com/zeiss/service-lens/cmd"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		panic(err)
	}
}
