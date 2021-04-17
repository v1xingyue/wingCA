package api

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wingCA/config"
	"wingCA/rootCA"
)

func confirmServiceSSL() {
	if _, err := os.Open(rootCA.SiteCertPath(config.Item.Wing.Domain)); err != nil {
		log.Printf(
			"may need this command : \n\n./wingCA issue --type site --email admin@ssl.wingca.com.cn  --common %s --site %s --site localhost \n\n",
			config.Item.Wing.Domain,
			config.Item.Wing.Domain,
		)
		log.Fatal(err)
	}
}

type caServiceHandler struct{}

func (cs *caServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func StartSSL() {
	var (
		err               error
		certBytes, kBytes []byte
		cert              tls.Certificate
	)
	confirmServiceSSL()

	certBytes, err = ioutil.ReadFile(rootCA.SiteCertPath(config.Item.Wing.Domain))
	if err != nil {
		log.Println(err)
	}

	kBytes, err = rootCA.UnWrapEncryptKey(rootCA.PrivateKeyPath(config.Item.Wing.Domain), config.Default.KeyPassword)

	cert, err = tls.X509KeyPair(certBytes, kBytes)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	server := http.Server{
		TLSConfig: tlsConfig,
		Addr:      ":443",
		Handler:   &caServiceHandler{},
	}

	err = server.ListenAndServeTLS("", "")

}
