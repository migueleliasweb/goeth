package main

import (
	"os"

	"github.com/migueleliasweb/goeth/goeth"
)

func main() {
	goeth.Run(os.Stdout)
}
