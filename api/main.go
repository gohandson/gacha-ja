package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.FormValue("q")
		items := strings.Split(q, ",")
		dist := make(map[string]int, len(items))
		var max int
		for _, item := range items {
			keyval := strings.Split(item, ":")
			if len(keyval) != 2 {
				http.Error(w, "Invalid query", http.StatusBadRequest)
				return
			}

			val, err := strconv.Atoi(keyval[1])
			if err != nil {
				http.Error(w, "Invalid query", http.StatusBadRequest)
				return
			}
			dist[keyval[0]] = val
			max += val
		}

		if max == 0 {
			http.Error(w, "Invalid query", http.StatusBadRequest)
			return
		}

		n := rand.Intn(max) + 1

		var count int
		for key, val := range dist {
			count += val
			if n <= count {
				fmt.Fprintln(w, key)
				return
			}
		}

		http.Error(w, "Invalid query", http.StatusBadRequest)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := net.JoinHostPort("", port)
	http.ListenAndServe(addr, nil)
}
