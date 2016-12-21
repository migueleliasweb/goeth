package main

import (
	"os"

	"github.com/migueleliasweb/goeth"
)

func main() {
	goeth.Run(os.Stdout)
}
