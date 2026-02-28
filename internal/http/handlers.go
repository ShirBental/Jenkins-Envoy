package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shirbental/jenkins-envoy/internal/jenkins"
)

type Server struct {
	client jenkins.Client
}

func NewServer(client jenkins.Client) *Server {
	return &Server{client: client}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.health)
	mux.HandleFunc("/jenkins/build", s.getBuild)
	return mux
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (s *Server) getBuild(w http.ResponseWriter, r *http.Request) {
	job := r.URL.Query().Get("job")
	numberStr := r.URL.Query().Get("number")

	if job == "" || numberStr == "" {
		http.Error(w, "missing job or number", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "invalid build number", http.StatusBadRequest)
		return
	}

	build, err := s.client.GetBuild(job, number)
	if err != nil {
		http.Error(w, "failed to fetch build", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(build)
}
