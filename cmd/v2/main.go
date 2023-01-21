package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
)

// https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application
// https://stackoverflow.com/a/40494806/17655004
func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("u8views.com", "dev.u8views.com"),
		Cache:      autocert.DirCache(os.Getenv("TLS_CERTIFICATES_DIR")),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12, // improves cert reputation score at https://www.ssllabs.com/ssltest/
		},
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("", "")) //Key and cert are coming from Let's Encrypt
}
