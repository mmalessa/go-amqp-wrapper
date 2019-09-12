package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strings"

	"github.com/streadway/amqp"
)

func getConnection(cfg ConfigConnection) (*amqp.Connection, error) {
	if strings.HasPrefix(cfg.Uri, "amqps:") {
		tlsCfg := new(tls.Config)
		tlsCfg.RootCAs = x509.NewCertPool()
		tlsCfg.ServerName = cfg.ServerName
		if ca, err := ioutil.ReadFile(cfg.SslCaCert); err == nil {
			tlsCfg.RootCAs.AppendCertsFromPEM(ca)
		}
		if cert, err := tls.LoadX509KeyPair(cfg.SslCert, cfg.SslKey); err == nil {
			tlsCfg.Certificates = append(tlsCfg.Certificates, cert)
		}
		return amqp.DialTLS(cfg.Uri, tlsCfg)
	} else {
		return amqp.Dial(cfg.Uri)
	}
}
