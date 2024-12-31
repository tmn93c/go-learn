package router

import (
	"encoding/json"
	"log"
	"my-project/internal/database"
	"net/http"

	"fmt"
	"time"

	"github.com/coder/websocket"
	"github.com/go-chi/chi/v5"
)

type indexResource struct {
	service database.Service
}

func NewIndexResource(service database.Service) *indexResource {
	return &indexResource{service: service}
}

func (s *indexResource) RegisterRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.helloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/websocket", s.websocketHandler)
	return r
}

// @Summary Get a hello word
// @Description Get a hello word message
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func (s *indexResource) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *indexResource) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.service.Health())
	_, _ = w.Write(jsonResp)
}

func (s *indexResource) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
