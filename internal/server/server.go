package server

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/u8views/go-u8views/internal/env"

	"golang.org/x/crypto/acme/autocert"
)

// https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application
// https://stackoverflow.com/a/40494806/17655004
func Run(handler http.Handler) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("u8views.com"),
		Cache:      autocert.DirCache(env.Must("TLS_CERTIFICATES_DIR")),
	}

	server := &http.Server{
		Addr:    ":https",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12, // improves cert reputation score at https://www.ssllabs.com/ssltest/
		},
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("", "")) //Key and cert are coming from Let's Encrypt
}
