package gomarathon

/*
 * Server Info
 * https://mesosphere.github.io/marathon/docs/rest-api.html#server-info
 */

// Get Info about the Marathon cluster
func (c *Client) Info() (*Response, error) {
	options := &RequestOptions{
		Path: "info",
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Leader returns the current leader
func (c *Client) Leader() (*Response, error) {
	options := &RequestOptions{
		Path: "leader",
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// DeleteLeader asks the current leader to abdicate
func (c *Client) DeleteLeader() (*Response, error) {
	options := &RequestOptions{
		Path:   "leader",
		Method: "DELETE",
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}
