package main

import (
	"os"

	"./goeth"
)

func main() {
	goeth.Run(os.Stdout)
}
