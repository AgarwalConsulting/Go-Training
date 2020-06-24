package main

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func handleIndex(w http.ResponseWriter, req *http.Request) {
	msg := "Success"

	w.Write([]byte(msg))
}

func main() {
	var m *autocert.Manager

	m = &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		HostPolicy: func(ctx context.Context, host string) error {
			if host != "ssl-server.demo.algogrit.com" {
				return errors.New("acme/autocert: host name " + host + " is not allowed")
			}
			return nil
		},
		Cache: autocert.DirCache("."),
	}

	mux := http.ServeMux{}

	mux.HandleFunc("/", handleIndex)

	s := http.Server{
		Handler:   &mux,
		Addr:      ":443",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}

	s.ListenAndServeTLS("", "")
}
