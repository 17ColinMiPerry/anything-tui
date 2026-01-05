package api

import "time"

// Workspace represents the workspace db object
type Workspace struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"createdAt"`
	Threads   []Thread  `json:"threads,omitempty"`
}

type WorkspaceResponse struct {
	Workspace *Workspace
}

type WorkspacesResponse struct {
	Workspaces []Workspace
}

// Thread represents the threads of a particular workspace
type Thread struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	WorkspaceID int       `json:"workspace_id"`
	CreatedAt   time.Time `json:"createdAt"`
	Workspace   Workspace `json:"workspace"`
}

type ThreadResponse struct {
	Thread *Thread
}

type ThreadsResponse struct {
	Threads []Thread
}

// APIError represents an error response from the API
type APIError struct {
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"-"`
}

func (e *APIError) ErrorMessage() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Error != "" {
		return e.Error
	}
	return "unknown error"
}
