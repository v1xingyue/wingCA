package rootCA

import (
	"bufio"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

// RevokeSiteCert 吊销一个服务端证书
// 只是在 revoke 列表中添加对应的revoke 记录
func RevokeSiteCert(certPath string) error {

	var (
		cert *x509.Certificate
		err  error
		f    *os.File
	)

	if cert, err = ParseCertificate(certPath); err != nil {
		return err
	}
	serial := cert.SerialNumber.Int64()

	revokeMessage := fmt.Sprintf("%d,%d,%s\n", serial, time.Now().Unix(), cert.Subject.CommonName)
	f, err = os.OpenFile(revokeListPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, newFileMode)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write([]byte(revokeMessage))
	return nil

}

// CrlBytes 生成对应的CRL 文件字节
// cat site.bytes | openssl crl -inform der -text
// 其中 序列号为 16进制
func CrlBytes() ([]byte, error) {
	rootCert, rootPriv, err := LoadCARootWithKey()
	if err != nil {
		return nil, err
	}
	revokeList := []pkix.RevokedCertificate{}

	revokeListFile, err := os.OpenFile(revokeListPath, os.O_RDONLY, newFileMode)
	if err != nil {
		return nil, err
	}
	bufReader := bufio.NewReader(revokeListFile)

	for {
		buf, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}
		items := strings.Split(string(buf), ",")
		if len(items) == 3 {
			serialInt, err := strconv.Atoi(items[0])
			revokeSeconds, err := strconv.Atoi(items[1])
			if err == nil {

				revokeTime := time.Unix(int64(revokeSeconds), 0)

				revokeList = append(revokeList, pkix.RevokedCertificate{
					SerialNumber:   big.NewInt(int64(serialInt)),
					RevocationTime: revokeTime,
				})
			}
		}
	}

	return rootCert.CreateCRL(rand.Reader, rootPriv, revokeList, time.Now(), time.Now().Add(crlLifetime))
}
