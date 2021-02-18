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
		if r.FormValue("delay") != "" {
			time.Sleep(delay())
		}
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
				fmt.Fprint(w, key)
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

func random(dist []int) int {
	if len(dist) == 0 {
		return -1
	}

	var sum int
	for _, d := range dist {
		sum += d
	}

	r := rand.Intn(sum)
	for i, d := range dist {
		r -= d
		if r <= 0 {
			return i
		}
	}

	return dist[len(dist)-1]
}

func delay() time.Duration {
	return []time.Duration{
		0,
		100 * time.Millisecond,
		1 * time.Second,
		10 * time.Second,
		30 * time.Second,
		60 * time.Second,
	}[random([]int{20, 20, 20, 20, 20})]
}
