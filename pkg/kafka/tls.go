package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
)

type TLS struct {
	Enable     bool   `json:"enable" yaml:"enable"`
	CaFile     string `json:"ca_file" yaml:"ca_file"`
	KeyFile    string `json:"key_file" yaml:"key_file"`
	CertFile   string `json:"cert_file" yaml:"cert_file"`
	SkipVerify bool   `json:"skip_verify" yaml:"skip_verify"`
}

func createTlsConfig(c TLS) (t *tls.Config) {
	t = &tls.Config{
		InsecureSkipVerify: c.SkipVerify,
	}
	if c.CertFile != "" && c.KeyFile != "" && c.CaFile != "" {
		cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
		if err != nil {
			log.Fatal(err)
		}

		caCert, err := os.ReadFile(c.CaFile)
		if err != nil {
			log.Fatal(err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		t = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: c.SkipVerify,
		}
	}
	return t
}
