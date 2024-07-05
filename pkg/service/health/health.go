package health

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
)

type healthResponse struct {
	Timestamp int64             `json:"timestamp"`
	State     map[string]string `json:"state"`
}

type getStatusCallback func() map[string]bool

func Start(address string, getStatus getStatusCallback) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler(getStatus))

	err := http.ListenAndServe(address, mux)
	if err != nil {
		// TODO
		logger.Error("Health service start failed: ", err)
	}
}

func healthHandler(getStatus getStatusCallback) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := getStatus()
		resp := healthResponse{
			Timestamp: time.Now().Unix(),
			State:     map[string]string{},
		}
		httpCode := http.StatusOK
		for name, ok := range status {
			if ok {
				resp.State[name] = "ok"
			} else {
				resp.State[name] = "error"
				httpCode = http.StatusBadRequest
			}
		}
		data, _ := json.Marshal(resp)

		writeResponse(w, httpCode, data)
	}
}

func writeResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
