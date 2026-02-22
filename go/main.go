package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func main() {
	s := &server{counts: make(map[string]int64)}
	http.HandleFunc("/api/view", s.cors(s.handleView))
	http.HandleFunc("/api/stats", s.cors(s.handleStats))
	log.Println("listening on :8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("listening on :", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type server struct {
	mu     sync.RWMutex
	counts map[string]int64
}

func (s *server) cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func normalizePath(raw string) string {
	if raw == "" {
		return "/"
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "/"
	}
	p := u.Path
	if p == "" {
		return "/"
	}
	return p
}

func (s *server) handleView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := normalizePath(r.URL.Query().Get("path"))
	s.mu.Lock()
	s.counts[path]++
	count := s.counts[path]
	s.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"path": path, "count": count})
}

func (s *server) handleStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	s.mu.RLock()
	out := make(map[string]int64, len(s.counts))
	for k, v := range s.counts {
		out[k] = v
	}
	s.mu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}
