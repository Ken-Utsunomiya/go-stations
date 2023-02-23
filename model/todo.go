package model

import "time"

type (
	// A TODO expresses basic struct for TODO model
	TODO struct {
		ID          int64     `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// A CreateTODORequest expresses request body for POST /todos
	CreateTODORequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	// A CreateTODOResponse expresses response body for POST /todos
	CreateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	// A ReadTODORequest expresses request body for GET /todos
	ReadTODORequest struct {
		PrevID int64 `json:"prev_id"`
		Size   int64 `json:"size"`
	}
	// A ReadTODOResponse expresses response body for GET /todos
	ReadTODOResponse struct {
		TODOs []*TODO `json:"todos"`
	}

	// A UpdateTODORequest expresses request body for PUT /todos
	UpdateTODORequest struct {
		ID          int64  `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	// A UpdateTODOResponse expresses response body for PUT /todos
	UpdateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	// A DeleteTODORequest expresses request body for DELETE /todos
	DeleteTODORequest struct {
		IDs []int64 `json:"ids"`
	}
	// A DeleteTODOResponse expresses response body for DELETE /todos
	DeleteTODOResponse struct{}
)
