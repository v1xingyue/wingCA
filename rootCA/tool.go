package rootCA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"software.sslmate.com/src/go-pkcs12"
)

// LoadCARoot 加载 CA 根证书
func LoadCARoot() (*x509.Certificate, error) {
	return ParseCertificate(rootCACertPath)
}

// LoadCARootWithKey 加载 CA 根证书 和 私钥
func LoadCARootWithKey() (*x509.Certificate, *rsa.PrivateKey, error) {

	var (
		rootCrt *x509.Certificate
		rootKey *rsa.PrivateKey
		err     error
	)

	if rootCrt, err = ParseCertificate(rootCACertPath); err != nil {
		return nil, nil, err
	}
	if rootKey, err = ParseKey(rootCAKeyPath, ""); err != nil {
		return rootCrt, nil, err
	}

	return rootCrt, rootKey, nil
}

// ParseCertificate 解析 证书
func ParseCertificate(path string) (*x509.Certificate, error) {

	var (
		certPEMBlock []byte
		err          error
		certDERBlock *pem.Block
		x509Cert     *x509.Certificate
	)

	if certPEMBlock, err = ioutil.ReadFile(path); err != nil {
		return nil, err
	}

	if certDERBlock, _ = pem.Decode(certPEMBlock); certDERBlock == nil {
		return nil, errors.New("pem file decode failed")
	}

	if x509Cert, err = x509.ParseCertificate(certDERBlock.Bytes); err != nil {
		return nil, err
	}

	return x509Cert, nil
}

// ParseKey 解析私钥文件
func ParseKey(path string, password string) (*rsa.PrivateKey, error) {

	var (
		keyBytes     []byte
		err          error
		privPem      *pem.Block
		privPemBytes []byte
	)

	if keyBytes, err = ioutil.ReadFile(path); err != nil {
		return nil, err
	}

	if privPem, _ = pem.Decode(keyBytes); privPem == nil {
		return nil, errors.New("pem file decode failed")
	}

	if password == "" {
		privPemBytes = privPem.Bytes
	} else {
		privPemBytes, err = x509.DecryptPEMBlock(privPem, []byte(password))
	}

	return x509.ParsePKCS1PrivateKey(privPemBytes)
}

// MakePKCS12 生成 客户端通用的 p12 证书
// openssl pkcs12 -export -clcerts -in ssl/client/client.crt -inkey ssl/private/client.key -out ssl/p12/client.p12
func MakePKCS12(certPath, keyPath, password string) error {

	var (
		err        error
		privateKey *rsa.PrivateKey
		cert       *x509.Certificate
	)

	if privateKey, err = ParseKey(keyPath, ""); err != nil {
		return err
	}

	if cert, err = ParseCertificate(certPath); err != nil {
		return err
	}

	rootCACerts, err := ParseCertificate(rootCACertPath)
	if err != nil {
		return err
	}

	pkbytes, err := pkcs12.Encode(rand.Reader, privateKey, cert, []*x509.Certificate{rootCACerts}, password)

	realCommonName := strings.Split(cert.Subject.CommonName[7:], "@")[0]

	err = ioutil.WriteFile(P12Path(realCommonName), pkbytes, newFileMode)
	return err
}

// InitDir 初始化文件夹结构
func InitDir() {

	os.Mkdir(RootCAPath, 0700)

	dirList := []string{
		RootCAPath + "/private",
		RootCAPath + "/client",
		RootCAPath + "/site",
		RootCAPath + "/root",
		RootCAPath + "/p12",
	}
	for _, d := range dirList {
		os.Mkdir(d, 0700)
	}
	ioutil.WriteFile(serialFile, []byte("10001"), 0700)
}

// SerialNumber 返回当前的序列值，并把文件内的值的 +1
func SerialNumber() *big.Int {
	if bs, err := ioutil.ReadFile(serialFile); err == nil {
		n := string(bs)
		if v, err := strconv.Atoi(n); err == nil {
			newSerial := fmt.Sprintf("%d", v+1)
			err = ioutil.WriteFile(serialFile, []byte(newSerial), newFileMode)
			if err != nil {
				log.Println("write serial file error : ", err)
			}
			log.Println("current crt serial number : ", v)
			return big.NewInt(int64(v))
		}
		log.Println(n)
	}
	return nil
}
