package rootCA

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

// SiteCertPath 站点证书路径
func SiteCertPath(commonName string) string {
	return fmt.Sprintf("%s/site/%s.crt", RootCAPath, commonName)
}

// PrivateKeyPath 私钥文件路径
func PrivateKeyPath(commonName string) string {
	return fmt.Sprintf("%s/private/%s.key", RootCAPath, commonName)
}

// ClientCertPath 返回客户端证书路径
func ClientCertPath(commonName string) string {
	return fmt.Sprintf("%s/client/%s.crt", RootCAPath, commonName)
}

// P12Path 返回P12 的路径
func P12Path(commonName string) string {
	return fmt.Sprintf("%s/p12/%s.p12", RootCAPath, commonName)
}

//IssueSite 使用自有CA 签发一个证书
// 返回证书 key 的字节
func IssueSite(commonName string, alternateIPs []net.IP, alternateDNS []string, email string) error {

	var (
		err error
	)
	rootCA, err := LoadCARoot()

	rootCAKey, err := ParseKey(rootCAKeyPath, rootCAKeyPassword)

	if err != nil {
		return err
	}
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: SerialNumber(),
		Subject: pkix.Name{
			CommonName: fmt.Sprintf("%s@%d", commonName, time.Now().Unix()),
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(defaultCertLifetime),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		// CRLDistributionPoints: []string{"http://localhost/crl"},
		EmailAddresses: []string{email},
	}

	template.IPAddresses = append(template.IPAddresses, alternateIPs...)
	template.DNSNames = append(template.DNSNames, alternateDNS...)

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, rootCA, &priv.PublicKey, rootCAKey)
	if err != nil {
		return err
	}

	// Generate cert
	certBuffer := bytes.Buffer{}
	if err := pem.Encode(&certBuffer, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return err
	}

	// Generate key
	keyBuffer := bytes.Buffer{}
	if err := pem.Encode(&keyBuffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}); err != nil {
		return err
	}

	err = ioutil.WriteFile(SiteCertPath(commonName), certBuffer.Bytes(), newFileMode)
	err = ioutil.WriteFile(PrivateKeyPath(commonName), keyBuffer.Bytes(), newFileMode)

	return err
}

// IssueClient 签发一对客户端证书
func IssueClient(clientName, email string) error {
	var (
		err error
	)
	rootCA, err := ParseCertificate(rootCACertPath)

	rootCAKey, err := ParseKey(rootCAKeyPath, rootCAKeyPassword)

	if err != nil {
		return err
	}
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: SerialNumber(),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(defaultClientCertLifetime),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		Subject: pkix.Name{
			CommonName: fmt.Sprintf("client-%s@%d", clientName, time.Now().Unix()),
		},
		EmailAddresses: []string{email},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, rootCA, &priv.PublicKey, rootCAKey)
	if err != nil {
		return err
	}

	// Generate cert
	certBuffer := bytes.Buffer{}
	if err := pem.Encode(&certBuffer, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return err
	}

	// Generate key
	keyBuffer := bytes.Buffer{}
	if err := pem.Encode(&keyBuffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}); err != nil {
		return err
	}

	err = ioutil.WriteFile(ClientCertPath(clientName), certBuffer.Bytes(), newFileMode)
	err = ioutil.WriteFile(PrivateKeyPath(clientName), keyBuffer.Bytes(), newFileMode)

	return err
}
