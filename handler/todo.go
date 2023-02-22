package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// ServeHTTP implements http.Handler interface.
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		prevID, _ := strconv.ParseInt(r.URL.Query().Get("prev_id"), 10, 64)
		size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
		req := model.ReadTODORequest{
			PrevID: prevID,
			Size:   size,
		}

		ctx := r.Context()
		res, err := h.Read(ctx, &req)
		if err != nil {
			fmt.Println("Error")
			return
		}

		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			fmt.Println("Error")
			return
		}
	case "POST":
		var ctr model.CreateTODORequest
		err := json.NewDecoder(r.Body).Decode(&ctr)
		if err != nil {
			fmt.Println("Error")
			return
		}
		if ctr.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		res, err := h.Create(ctx, &ctr)
		if err != nil {
			fmt.Println("Error")
			return
		}

		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			fmt.Println("Error")
			return
		}
	case "PUT":
		var req model.UpdateTODORequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("Error")
			return
		}
		if req.ID == 0 || req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		res, err := h.Update(ctx, &req)
		if err != nil {
			if err.Error() == "Not Found" {
				w.WriteHeader(http.StatusNotFound)
			}
			fmt.Println("Error")
			return
		}

		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			fmt.Println("Error")
			return
		}
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	res := &model.CreateTODOResponse{TODO: *todo}
	return res, err
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	defaultSize := int64(5)
	size := req.Size
	if size == 0 {
		size = defaultSize
	}

	todos, err := h.svc.ReadTODO(ctx, req.PrevID, size)
	if err != nil {
		return nil, err
	}
	res := &model.ReadTODOResponse{TODOs: todos}
	return res, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	todo, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}
	res := &model.UpdateTODOResponse{TODO: *todo}
	return res, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
