package main

import (
	"io"
	"net/http"
	"strings"
)

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}
func copyHeader(dst, src http.Header) {
	for k, vs := range src {
		switch strings.ToLower(k) {
		case /*"connection",*/ "keepalive", "proxy-authenticate, proxy-authorization, te, trailer, transfer-encoding":
			// skip
			vs = []string{}
		}
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}
