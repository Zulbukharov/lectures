package delivery

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"post/internal/app/adding"
	"post/internal/models"

	"github.com/gorilla/mux"
)

const (
	jsonContentType = "application/json"
	ADDED           = "post successfully created"
)

type Server struct {
	router *mux.Router
	adder  adding.Service
}

func NewServer(addingService adding.Service) *Server {
	s := &Server{
		router: mux.NewRouter(),
		adder:  addingService,
	}
	s.router.HandleFunc("/post", respHandler(s.AddPost)).Methods("POST")
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func respHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = Response{Error: err.Error()}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}

func (s *Server) AddPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var post models.PostInsert
	var res Response

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		return res, http.StatusBadRequest, errors.New("failed to parse post")
	}
	err := s.adder.AddPost(post)
	if err != nil {
		return res, http.StatusBadRequest, errors.New("failed to parse post")
	}
	res.Message = ADDED
	return res, http.StatusOK, nil
}
