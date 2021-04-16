package rootCA

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"time"
)

// CertExpired 证书是否过期
func CertExpired(commonName string, certType int) bool {
	var (
		cert *x509.Certificate
		err  error
	)
	if certType == CertTypeClient {
		cert, err = ParseCertificate(ClientCertPath(commonName))
	} else if certType == CertTypeSite {
		cert, err = ParseCertificate(SiteCertPath(commonName))
	}
	if err == nil {
		return cert.NotAfter.Before(time.Now())
	}
	return false
}

func UnWrapEncryptKey(kfile, pass string) ([]byte, error) {
	var (
		keyBytes, privPemBytes []byte
		err                    error
		privPem                *pem.Block
		priKey                 *rsa.PrivateKey
	)
	keyBytes, err = ioutil.ReadFile(kfile)
	if err == nil {
		if privPem, _ = pem.Decode(keyBytes); privPem != nil {
			privPemBytes, err = x509.DecryptPEMBlock(privPem, []byte(pass))
			priKey, err = x509.ParsePKCS1PrivateKey(privPemBytes)
			keyBuffer := bytes.Buffer{}
			block := &pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(priKey),
			}
			pem.Encode(&keyBuffer, block)
			return keyBuffer.Bytes(), nil
		}
	}
	return nil, err
}

func appendMiddleCrt(filename string) {
	middleBytes, err := ioutil.ReadFile(middleCACertPath)
	if err == nil {
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		defer f.Close()
		if err == nil {
			f.Write(middleBytes)
		}
	}
}
