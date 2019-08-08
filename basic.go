package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type User struct {
	UserId   string
	Password string
}

func basicAuth(w http.ResponseWriter, r *http.Request, users []User) bool {

	var auth = r.Header.Get("Proxy-Authorization")

	if ms := strings.Split(auth, " "); len(ms) == 2 && ms[0] == "Basic" {

		// check user:password
		up, err := base64.StdEncoding.DecodeString(ms[1])

		if err == nil {
			if ms := strings.Split(string(up), ":"); len(ms) == 2 {

				var user, password = ms[0], ms[1]
				var ok = false

				for _, u := range users {
					if u.UserId == user && u.Password == password {
						ok = true
						break
					}
				}

				if ok {
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
