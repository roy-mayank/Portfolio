package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var activeUsers = make(map[string]time.Time)
var mu sync.Mutex

func handleHeartbeat(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ") 

	parts := strings.Split(token, ".")
	if len(parts) >= 2 {
		payload, _ := base64.RawURLEncoding.DecodeString(parts[1])
		log.Printf("JWT Payload size: %d bytes", len(payload))
	}

	mu.Lock()
	activeUsers["user_id_here"] = time.Now()
	mu.Unlock()
}