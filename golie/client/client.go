/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/

package golie

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func Client() {
	rolieserver := "http://localhost:3000/feed"
	caFile := "../examples/certs/cert.pem"
	certFile := "../examples/certs/cert.pem"
	keyFile := "../examples/certs/key.pem"
	timeout := 15

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	// Make a request
	r, err := client.Get(rolieserver)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer r.Body.Close()

	if r.StatusCode == 200 {
		fmt.Println("Response status:", r.Status)
	}

	buf, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(buf))

}
