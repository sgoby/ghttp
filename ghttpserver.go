package ghttp

import (
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"crypto/x509"
)

func (g *GHttp) getHttpServer(addr string) (*http.Server, error) {
	return &http.Server{
		Addr:    addr,
		Handler: g,
	},nil
}
func (g *GHttp) getHttpServerTLS(addr string,caCertPath string) (*http.Server, error) {
	pool := x509.NewCertPool()
	//caCertPath := "ca.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}
	//
	pool.AppendCertsFromPEM(caCrt)
	return &http.Server{
		Addr:    addr,
		Handler: g,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}, nil
}
