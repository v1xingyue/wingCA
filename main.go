package main

import (
	"wingCA/cmd"
	_ "wingCA/config"
	"wingCA/rootCA"
)

func main() {
	rootCA.CrlBytes()
	cmd.Main()
}
