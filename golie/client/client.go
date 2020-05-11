/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/

package golie

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rolieup/golie/internal/xml"
	"github.com/rolieup/golie/pkg/api"
	log "github.com/sirupsen/logrus"
)

func Client() {
	rolieserver := "http://localhost:3000/"
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
				MinVersion:   tls.VersionTLS12,
			},
		},
	}

	// Make a request
	req, err := http.NewRequest("GET", rolieserver, nil)
	//req.Header.Set("Accept", "atomsvc+xml")
	//req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	if r.StatusCode == 200 {
		fmt.Println("Response status:", r.Status)
	}

	servedContent := r.Header.Get("Content-Type")
	fmt.Println(servedContent)

	switch {
	case strings.Contains(servedContent, "atomsvc+xml"):
		serviceD := &api.Service{}
		serviceD.Xmlns = "https://www.w3.org/2005/Atom"

		err = xml.NewDecoder(r.Body).Decode(&serviceD)
		if err != nil {
			fmt.Printf("Failed to decode the xml service document: %s", err)
		}

		for i, w := range serviceD.Workspaces {
			i = i + 1
			fmt.Println("Document Type: Workspace")
			fmt.Printf("Title: %s\n", w.Title)
			fmt.Println("Collections:")
			for _, c := range w.Collections {
				fmt.Printf("\tTitle: %s\n", c.Title)
				fmt.Printf("\t\tDocument URL: %s\n", c.Href)
				fmt.Printf("\t\tCategory Type: %s\n", c.Categories.Category[0].Term)
				fmt.Printf("\t\tCategory Scheme: %s\n", c.Categories.Category[0].Scheme)
				fmt.Printf("\t\tService Information:\n")
				fmt.Printf("\t\t  Type: %s\n", c.Link.Rel)
				fmt.Printf("\t\t   URL: %s\n", c.Link.Href)
			}
			fmt.Println()
		}
	case strings.Contains(servedContent, "json"):
		serviceD := &api.JSONServiceRoot{}
		err := json.NewDecoder(r.Body).Decode(&serviceD)
		if err != nil {
			fmt.Printf("Failed to decode the json service document: %s", err)
		}

		for i, w := range serviceD.Service.Workspaces {
			i = i + 1
			fmt.Println("Document Type: Workspace")
			fmt.Printf("Title: %s\n", w.Title)
			fmt.Println("Collections:")
			for _, c := range w.Collections {
				fmt.Printf("\tTitle: %s\n", c.Title)
				fmt.Printf("\t\tDocument URL: %s\n", c.Href)
				fmt.Printf("\t\tCategory Type: %s\n", c.Categories.Category[0].Term)
				fmt.Printf("\t\tCategory Scheme: %s\n", c.Categories.Category[0].Scheme)
				fmt.Printf("\t\tService Information:\n")
				fmt.Printf("\t\t  Type: %s\n", c.Link.Rel)
				fmt.Printf("\t\t   URL: %s\n", c.Link.Href)
			}
			fmt.Println()
		}
	case strings.Contains(servedContent, "atom+xml"):
		buf, _ := ioutil.ReadAll(r.Body)

		data, err := xml.MarshalIndent(buf, "", "  ")
		if err != nil {
			fmt.Printf("Failed to marshal the feed: %s", err)
			os.Exit(1)
		}
		fmt.Println(string(data))
	}
	//buf, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(buf))

}
