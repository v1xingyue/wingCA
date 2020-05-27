package rootCA

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SampleWeb 启动一个 实例 站点
func SampleWeb(certPath, keyPath string) {
	go serveCrl()
	log.Println("Sample site started ")

	s := &http.Server{
		Addr: ":443",
		Handler: &myhandler{
			content: "sample https site",
		},
	}

	if e := s.ListenAndServeTLS(certPath, keyPath); e != nil {
		log.Fatal("ListenAndServe: ", e)
	}
}

// SampleDoubleWeb 启动一个双向认证的站点
// curl -k --cert client.pem --key key.pem https://www.xxxx.com
func SampleDoubleWeb(certPath, keyPath, rootCAPath string) {
	go serveCrl()
	log.Println(" Start Double validate site ...")
	pool := x509.NewCertPool()
	caCertPath := rootCAPath

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err: ", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr: ":443",
		Handler: &myhandler{
			content: "hello double validate https",
		},
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
	content string
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.content)
}

// cat site.crl | openssl crl -inform der -text
func serveCrl() {
	http.HandleFunc("/crl", func(w http.ResponseWriter, req *http.Request) {
		log.Println("crl request come ", req.RemoteAddr)
		w.Header().Add("Content-Disposition", "attachment; filename=site.crl")
		crlBytes, err := CrlBytes()
		if err != nil {
			log.Println(err)
		}
		w.Write(crlBytes)
	})
	http.ListenAndServe(":80", nil)
}
