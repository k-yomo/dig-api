package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main()  {
	http.HandleFunc("/lookup/ip", ipLookupHandler)

	port := "1323"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}

func ipLookupHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'domain' must be set in query parameter"))
		return
	}

	ips, err := net.LookupIP(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if len(ips) == 0 {
		w.Write([]byte("no record"))
	}

	var res string
	for _, ip := range ips {
		res += fmt.Sprintf("%s\n", ip.String())
	}
	w.Write([]byte(res))
}
