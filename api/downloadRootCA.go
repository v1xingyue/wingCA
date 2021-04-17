package api

import (
	"fmt"
	"log"
	"net/http"
	"wingCA/config"
	"wingCA/rootCA"
)

type basicHTTPServer struct {
}

func (bh *basicHTTPServer) Start() {
	go func() {
		basicConfig := config.Item.BasicConfig
		bindAddr := fmt.Sprintf("%s:%d", basicConfig.BindHost, basicConfig.BindPort)
		log.Println("Start rootca download service at  : ", bindAddr)

		log.Printf("You need download root CA from : \n\n")

		log.Printf("http://%s/download/rootca\n", bindAddr)

		http.HandleFunc("/download/rootca", downloadFileHandler(rootCA.RootCACertPath, "rootCASaved.crt"))
		http.ListenAndServe(bindAddr, nil)
	}()
}

func StartDownloadRootCA() {
	bh := basicHTTPServer{}
	bh.Start()
}
