package main

import (
	"github.com/zeiss/service-lens/cmd/server/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		panic(err)
	}
}
