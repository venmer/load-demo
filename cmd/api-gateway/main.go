package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type ProxyClient struct {
	userServiceURL string
	client         *http.Client
}

func NewProxyClient(userServiceURL string) *ProxyClient {
	return &ProxyClient{
		userServiceURL: userServiceURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *ProxyClient) ProxyRequest(w http.ResponseWriter, r *http.Request, path string) {
	url := p.userServiceURL + path

	// Создаем новый запрос
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Копируем заголовки
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Выполняем запрос
	resp, err := p.client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Копируем статус и заголовки ответа
	w.WriteHeader(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Копируем тело ответа
	io.Copy(w, resp.Body)
}

func main() {
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		userServiceURL = "http://localhost:8081"
	}

	proxy := NewProxyClient(userServiceURL)

	r := mux.NewRouter()
	
	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "api-gateway"})
	}).Methods("GET")

	// Проксируем запросы к user-service
	r.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		proxy.ProxyRequest(w, r, "/api/users")
	}).Methods("GET", "POST")

	r.HandleFunc("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		path := "/api/users/" + mux.Vars(r)["id"]
		proxy.ProxyRequest(w, r, path)
	}).Methods("GET", "PUT", "DELETE")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API Gateway starting on port %s", port)
	log.Printf("Proxying to User Service at %s", userServiceURL)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

