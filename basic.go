package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func basicAuth(w http.ResponseWriter, r *http.Request) bool {
	var auth = r.Header.Get("Proxy-Authorization")
	if ms := strings.Split(auth, " "); len(ms) == 2 && ms[0] == "Basic" {
		// 验证用户
		up, err := base64.StdEncoding.DecodeString(ms[1])
		log.Printf("%s", up)
		if err == nil {
			if ms := strings.Split(string(up), ":"); len(ms) == 2 {
				if ms[0] == "test" {
					return true
				}
			}
		}
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.WriteHeader(http.StatusProxyAuthRequired)
		w.Header().Set("Proxy-Authenticate", `Basic realm="Http Proxy"`)
	}

	return false
}
