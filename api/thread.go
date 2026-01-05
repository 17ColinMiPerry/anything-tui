package api

import "fmt"

func (c *Client) ListThreads(workspaceSlug string) ([]Thread, error) {
	workspace, err := c.GetWorkspace(workspaceSlug)
	if err != nil {
		return nil, err
	}

	var threads []Thread

	for _, th := range workspace.Threads {
		currThread, err := c.GetThread(workspaceSlug, th.Slug)
		if err != nil {
			threads = append(threads, th)
		} else {
			threads = append(threads, *currThread)
		}
	}

	return threads, nil
}

func (c *Client) GetThread(workspaceSlug string, threadSlug string) (*Thread, error) {
	var threadResponse ThreadResponse
	url := fmt.Sprintf("/workspace/%s/thread/%s/update", workspaceSlug, threadSlug)
	err := c.post(url, map[string]string{}, &threadResponse)
	if err != nil {
		return nil, err
	}

	return threadResponse.Thread, err
}
