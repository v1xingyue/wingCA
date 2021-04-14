package main

import (
	"log"
	"wingCA/cmd"
	_ "wingCA/config"
)

var (
	// ComileTime 设置软件版本
	ComileTime = ""
)

func main() {
	log.Println(ComileTime)
	// rootCA.CrlBytes()
	cmd.Main()
}
