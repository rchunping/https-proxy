package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"strings"
)

func main() {
	var pemPath, keyPath, proto, listen, users string
	flag.StringVar(&pemPath, "pem", "server.pem", "path to pem file")
	flag.StringVar(&keyPath, "key", "server.key", "path to key file")
	flag.StringVar(&proto, "proto", "http", "Proxy protocol (http or https)")
	flag.StringVar(&listen, "listen", ":8080", "listen address, default :8080")
	flag.StringVar(&users, "users", "", "user:password list")
	flag.Parse()
	if proto != "http" && proto != "https" {
		log.Fatal("Protocol must be either http or https")
	}

	var userList []User
	for _, up := range strings.Split(users, ";") {
		if ms := strings.Split(up, ":"); len(ms) == 2 && len(ms[0]) > 0 {
			userList = append(userList, User{ms[0], ms[1]})
		}
	}

	server := &http.Server{
		Addr: listen,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if len(userList) > 0 && !basicAuth(w, r, userList) {
				return
			}

			if r.Method == http.MethodConnect {
				handleTunneling(w, r)
			} else {
				handleHTTP(w, r)
			}
		}),
		// TLSNextProto not-nil to disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	if proto == "http" {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	}
}
