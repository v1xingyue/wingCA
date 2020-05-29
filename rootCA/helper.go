package rootCA

import (
	"crypto/x509"
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
