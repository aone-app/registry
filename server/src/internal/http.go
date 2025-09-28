package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func StartHttpServer(mux *http.ServeMux, reg *Registry) {
	mux.HandleFunc("/api/registry/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var body struct {
			Service string `json:"service"`
			Address string `json:"address"`
			TTL     int    `json:"ttl"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := reg.Register(body.Service, body.Address, body.TTL); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/registry/deregister", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var body struct {
			Service string `json:"service"`
			Address string `json:"address"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := reg.Unregister(body.Service, body.Address); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/registry/list", func(w http.ResponseWriter, r *http.Request) {
		serviceName := r.URL.Query().Get("service")
		nodes, err := reg.GetNodes(serviceName)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"nodes": nodes})
	})
}

func StartHttpServer1(reg *Registry) {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received register request: %v", r.Body)

		var req struct {
			Service string `json:"service"`
			Addr    string `json:"addr"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if err := reg.Register(req.Service, req.Addr, 60); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/unregister", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Service string `json:"service"`
			Addr    string `json:"addr"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if err := reg.Unregister(req.Service, req.Addr); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/nodes", func(w http.ResponseWriter, r *http.Request) {
		service := r.URL.Query().Get("service")
		nodes, err := reg.GetNodes(service)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(nodes)
	})

	//http.ListenAndServe(":8080", nil)
}
