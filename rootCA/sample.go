package rootCA

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// SampleWeb 启动一个 实例 站点
func SampleWeb(certPath, keyPath string) {
	fmt.Println("..")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	})
	if e := http.ListenAndServeTLS(":443", certPath, keyPath, nil); e != nil {
		log.Fatal("ListenAndServe: ", e)
	}
}
