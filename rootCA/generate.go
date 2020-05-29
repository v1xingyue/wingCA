package rootCA

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"time"
)

// InitRootCA 初始化 根 CA 证书
func InitRootCA(pkiName pkix.Name) error {
	var (
		err       error
		ca        *x509.Certificate
		caPrivKey *rsa.PrivateKey
		caBytes   []byte
	)

	// set up our CA certificate
	ca = &x509.Certificate{
		SerialNumber:          SerialNumber(),
		Subject:               pkiName,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// create our private and public key
	if caPrivKey, err = rsa.GenerateKey(rand.Reader, 4096); err != nil {
		return err
	}

	// create the CA
	if caBytes, err = x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey); err != nil {
		return err
	}

	// pem encode
	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	ioutil.WriteFile(rootCACertPath, caPEM.Bytes(), 0755)

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})

	ioutil.WriteFile(rootCAKeyPath, caPrivKeyPEM.Bytes(), 0700)
	return nil
}
