package main

import (
	"log"
	"net/http"

	registryconnect "nerosoft.com/aone/registry-server/gen/registryconnect"
	internal "nerosoft.com/aone/registry-server/internal"
)

func main() {
	reg := internal.NewRegistry()
	service := internal.NewRegistryService(reg)

	mux := http.NewServeMux()

	path, handler := registryconnect.NewRegistryServiceHandler(service)
	mux.Handle(path, handler)

	// mux.HandleFunc("/api/registry/register", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// 	var body struct {
	// 		Service string `json:"service"`
	// 		Address string `json:"address"`
	// 		TTL     int    `json:"ttl"`
	// 	}
	// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	// 	if err := reg.Register(body.Service, body.Address, body.TTL); err != nil {
	// 		http.Error(w, err.Error(), 500)
	// 		return
	// 	}
	// 	json.NewEncoder(w).Encode(map[string]any{"ok": true})
	// })

	// mux.HandleFunc("/api/registry/list", func(w http.ResponseWriter, r *http.Request) {
	// 	serviceName := r.URL.Query().Get("service")
	// 	nodes, err := reg.GetNodes(serviceName)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 500)
	// 		return
	// 	}
	// 	json.NewEncoder(w).Encode(map[string]any{"nodes": nodes})
	// })

	// 启动 HTTP
	internal.StartHttpServer(mux, reg)

	// 启动 RPC
	//go StartRPCServer(reg)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	select {} // 阻塞
}
