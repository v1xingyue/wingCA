package rootCA

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
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

// SampleDoubleWeb 启动一个双向认证的站点
func SampleDoubleWeb(certPath, keyPath, rootCAPath string) {
	fmt.Println(" Start Double validate site ...")
	pool := x509.NewCertPool()
	caCertPath := rootCAPath

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err: ", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":443",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		fmt.Println(err)
	}
}

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"hello, double validate site world!\n")
}
