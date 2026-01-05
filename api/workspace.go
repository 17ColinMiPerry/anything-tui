package api

import "fmt"

func (c *Client) ListWorkspaces() ([]Workspace, error) {
	var resp WorkspacesResponse
	err := c.get("/workspaces", &resp)
	if err != nil {
		return nil, err
	}
	return resp.Workspaces, nil
}

func (c *Client) GetWorkspace(slug string) (*Workspace, error) {
	var resp struct {
		Workspaces []Workspace `json:"workspace"`
	}

	err := c.get("/workspace/"+slug, &resp)
	if err != nil {
		return nil, err
	}
	if len(resp.Workspaces) == 0 {
		return nil, fmt.Errorf("no workspaces at slug %s", slug)
	}

	return &resp.Workspaces[0], nil
}
