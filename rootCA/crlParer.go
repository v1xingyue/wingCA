package rootCA

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"log"
	"net"
	"wingCA/config"
)

var (
	oidEmailAddress = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}
)

func parseCSRContent(buffer []byte) {
	var (
		csrBlock *pem.Block
		x509CSR  *x509.CertificateRequest
		err      error
	)
	if csrBlock, _ = pem.Decode(buffer); csrBlock != nil {
		if x509CSR, err = x509.ParseCertificateRequest(csrBlock.Bytes); err != nil {
			log.Println(err)
			return
		}
		log.Println("csr request ....")
		log.Println(x509CSR.Subject.CommonName)
		log.Println(x509CSR.PublicKey)
		log.Println(x509CSR.EmailAddresses)
		log.Println(x509CSR.DNSNames)
		log.Println(x509CSR.IPAddresses)
	}
}

// commonName 可以是 IP 或者 域名
func makeCSR(emailAddress, commonName string, names []string, addrs []net.IP) ([]byte, []byte) {

	var (
		err error
	)

	keyPri, _ := rsa.GenerateKey(rand.Reader, 1024)

	subj := pkix.Name{
		CommonName:         commonName,
		Country:            []string{config.Default.Country},
		Province:           []string{config.Default.Province},
		Locality:           []string{config.Default.Locality},
		Organization:       []string{config.Default.Org},
		OrganizationalUnit: []string{config.Default.OrgUnit},
	}
	rawSubj := subj.ToRDNSequence()
	rawSubj = append(rawSubj, []pkix.AttributeTypeAndValue{
		{Type: oidEmailAddress, Value: emailAddress},
	})

	asn1Subj, _ := asn1.Marshal(rawSubj)
	template := x509.CertificateRequest{
		RawSubject:         asn1Subj,
		EmailAddresses:     []string{emailAddress},
		SignatureAlgorithm: x509.SHA256WithRSA,
		DNSNames:           names,
		IPAddresses:        addrs,
	}

	csrBuffer := new(bytes.Buffer)
	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &template, keyPri)
	pem.Encode(csrBuffer, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	priKeyBuffer := bytes.Buffer{}
	block := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(keyPri)}

	block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(config.Default.KeyPassword), x509.PEMCipherAES256)
	if err != nil {
		return nil, nil
	}

	if err := pem.Encode(&priKeyBuffer, block); err != nil {
		return nil, nil
	}

	return csrBuffer.Bytes(), priKeyBuffer.Bytes()
}
