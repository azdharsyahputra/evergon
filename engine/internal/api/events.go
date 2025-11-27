package api

import (
	"encoding/json"
	"net/http"
	"time"

	"evergon/engine/internal/manager"
	"evergon/engine/internal/scanner"
	"evergon/engine/internal/util/resolver"
)

type ProjectStatusEvent struct {
	Project string `json:"project"`
	Running bool   `json:"running"`
	Port    string `json:"port"`
}

func RegisterSSE(mux *http.ServeMux, res *resolver.Resolver) {

	mux.HandleFunc("/events/project-status", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		for {
			projects := scanner.Scan(res)

			for _, p := range projects {
				cfg, _ := manager.LoadProjectConfig(p.Name)
				running := manager.IsProjectRunning(p.Name)

				ev := ProjectStatusEvent{
					Project: p.Name,
					Running: running,
					Port:    cfg.PHPPort,
				}

				jsonData, _ := json.Marshal(ev)

				w.Write([]byte("data: " + string(jsonData) + "\n\n"))
				flusher.Flush()
			}

			time.Sleep(2 * time.Second)
		}
	})
}
