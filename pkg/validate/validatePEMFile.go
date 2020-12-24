package validate

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// ValidatePEM will check for pem file validity by decoding the certificate
func ValidatePEM(f string) {
	certFile := f
	certPEM, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Error("Failed read certificatee file" + err.Error())
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(certPEM))
	if !ok {
		log.Error("Failed to parse certificate, file will be removed")
		err := os.Remove(certFile)
		if err != nil {
			log.Error("could not delete file" + err.Error())
			return
		}
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		log.Error("Failed to parse certificate PEM" + err.Error())
		return
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Error("Failed to parse certificate" + err.Error())
		//fmt.Errorf("Failed to parse certificate: %v", err.Error())
		return

	}

	opts := x509.VerifyOptions{
		// DNSName: name,
		Roots: roots,
		// accept all keyusages kind
		KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}

	if _, err := cert.Verify(opts); err != nil {
		log.Error("Failed to verify certificate" + err.Error())
		// fmt.Errorf("Failed to verify certificate: %v", err.Error())
		return
	}

	// print cert infos
	certIssuer := cert.Issuer
	certCN := cert.Subject.CommonName
	certEndDate := cert.NotAfter.String()

	log.Info("successfully validated certificate for issuer: [", certIssuer, "] Common Name: [", certCN, "] valid until [", certEndDate, "]")
	return

}
