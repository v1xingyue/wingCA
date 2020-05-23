package rootCA

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

// ParseCertificate 解析 证书
func ParseCertificate(path string) (*x509.Certificate, error) {
	certPEMBlock, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		return nil, errors.New("pem file decode failed")
	}
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return x509Cert, nil
}

// ParseKey 解析私钥文件
func ParseKey(path string, password string) (*rsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	privPem, _ := pem.Decode(keyBytes)
	var privPemBytes []byte
	if password == "" {
		privPemBytes = privPem.Bytes
	} else {
		privPemBytes, err = x509.DecryptPEMBlock(privPem, []byte(password))
	}

	return x509.ParsePKCS1PrivateKey(privPemBytes)
}
