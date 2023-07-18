package main

import (
	"fmt"
	"net/http"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server Test")
	}

	httpsServer := &http.Server{
		Addr:    ":443",
		Handler: http.HandlerFunc(handler),
	}

	httpServer := &http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
		}),
	}

	go func() {
		err := httpsServer.ListenAndServeTLS("ssl/cert.pem", "ssl/key.pem")
		if err != nil {
			fmt.Println("Error al iniciar el servidor HTTPS:", err)
		}
	}()

	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Println("Error al iniciar el servidor HTTP:", err)
	}
}
