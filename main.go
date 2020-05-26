package main

import (
	"wingCA/cmd"
	"wingCA/rootCA"
)

func main() {
	rootCA.CrlBytes()
	cmd.Main()
}
