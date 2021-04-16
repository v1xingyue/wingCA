package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"wingCA/rootCA"
)

func downloadCAHandler(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadFile(rootCA.RootCACertPath); err == nil {
		w.Header().Add("Content-Type", "application/force-download")
		w.Header().Add("Content-Disposition", "attachment;filename=rootCA.crt")
		w.Header().Add("Content-Transfer-Encoding", "binary")
		w.Write(data)
	} else {
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}
}

func StartDownloadRootCA() {
	go func() {
		log.Println("Start rootca download service at :9999 .")

		log.Printf("You need download root CA from : \n\n")

		addrs, err := net.InterfaceAddrs()
		if err == nil {
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					log.Printf("http://%s:9999/download/rootca\n", ipnet.IP.String())
				}

			}
		}

		http.HandleFunc("/download/rootca", downloadCAHandler)
		http.ListenAndServe(":9999", nil)
	}()
}
