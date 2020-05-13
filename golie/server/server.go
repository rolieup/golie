/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/

package golie

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

func Server(debug bool) {
	caFile := "../examples/certs/cert.pem"
	certFile := "../examples/certs/cert.pem"
	keyFile := "../examples/certs/key.pem"
	read_timeout := 5
	write_timeout := 10
	idle_timeout := 120
	host := "127.0.0.1"
	port := "3000"

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		MinVersion: tls.VersionTLS12,
	}
	tlsConfig.BuildNameToCertificate()

	log.Info("Starting the GOLIE server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", service)
	mux.HandleFunc("/service", service)
	mux.HandleFunc("/feed", feed)

	srv := &http.Server{
		Addr:         net.JoinHostPort(host, port),
		TLSConfig:    tlsConfig,
		Handler:      mux,
		ReadTimeout:  time.Duration(read_timeout) * time.Second,
		WriteTimeout: time.Duration(write_timeout) * time.Second,
		IdleTimeout:  time.Duration(idle_timeout) * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go serverShutdown(srv, idleConnsClosed)

	if debug {
		log.Warnf("Running the GOLIE server on host '%s' and listening on port '%s'", host, port)
		log.Warn("The server is running unencrypted! Never run unencrypted in production!")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalf("Error starting the GOLIE server: %v", err)
		}
	} else {
		log.Infof("Running the GOLIE server in TLS mode on host '%s' and listening on port '%s'", host, port)
		if err := srv.ListenAndServeTLS(certFile, keyFile); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalf("Error starting the GOLIE server: %v", err)
		}
	}
	<-idleConnsClosed
	log.Infof("GOLIE server '%s' has shutdown successfully!", host)
}

func serverShutdown(srv *http.Server, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	log.Info("Shutting down GOLIE server...")

	srv.SetKeepAlivesEnabled(false)

	// We received an interrupt signal, shut down.
	if err := srv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Fatalf("Error shutting down GOLIE server: %v", err)
	}
	close(idleConnsClosed)
}

func service(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("../examples/rolie/service/nvd.json")
	//fp := path.Join("../examples/rolie/service/nvd.xml")
	log.Infof("Received %s request from %s for ROLIE Service document at %s", r.Method, r.RemoteAddr, fp)
	acpt := r.Header["Accept"]
	switch acpt[0] {
	case "application/atomsvc+xml":
		log.Infof("Setting Content-Type: %s", acpt)
		w.Header().Set("Content-Type", "application/atomsvc+xml; charset=utf-8")
	case "application/json":
		log.Infof("Setting Content-Type: %s", acpt)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	default:
		w.Header().Set("Content-Type", "application/atomsvc+xml; charset=utf-8")
	}

	http.ServeFile(w, r, fp)
}

func feed(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("../examples/rolie/feed/gov.nist.nvd.cve.recent.xml")
	log.Infof("Received %s request from %s for ROLIE Feed at %s", r.Method, r.RemoteAddr, fp)
	switch acpt := r.Header["Accept"][0]; {
	case acpt == "application/atom+xml":
		log.Infof("Setting Content-Type: %s", acpt)
		w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	case acpt == "application/json":
		log.Infof("Setting Content-Type: %s", acpt)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	http.ServeFile(w, r, fp)
}

func entry(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("../examples/rolie/feed/gov.nist.nvd.cve.recent.xml")
	log.Infof("Received %s request from %s for ROLIE Feed at %s", r.Method, r.RemoteAddr, fp)
	w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, fp)
}
