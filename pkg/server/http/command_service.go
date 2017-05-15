package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

//createCommandResponse represents the http response send by the CreateCommand handler
type createCommandResponse struct {
	GroupID int64  `json:"group-id,omitempty"`
	Err     string `json:"error,omitsempty"`
}

//CreateCommand handler the command creation. It is called on route POST /cmd/.
//This func calls a sigstat.Client (typically a database client) and insert a new command
func (h *Handler) CreateCommand(w http.ResponseWriter, r *http.Request) {
	// Decode request.
	var cmd sigstat.Command

	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		log.Fatal(err)
	}

	//Create command
	id, errc := h.client.CommandService().CreateCommand(cmd)

	//Handle error from client
	if errc != nil {
		if err := json.NewEncoder(w).Encode(&createCommandResponse{Err: errc.Error()}); err != nil {
			log.Fatal(err)
		}
	}

	//Handle json serialization
	if err := json.NewEncoder(w).Encode(&createCommandResponse{GroupID: id}); err != nil {
		log.Fatal(err)
	}
}

//UpdateStatus handler PUT request that update the status for a given Command ID.
func (h *Handler) UpdateStatus(w http.ResponseWriter, r *http.Request) {

	cmd := sigstat.Command{
		Status: "running",
	}

	h.client.CommandService().UpdateStatus(cmd)
}
