package teamspeak

type ClientList struct {
	SessionID  string `json:"clid"`
	DatabaseID string `json:"client_database_id"`
	Nickname   string `json:"client_nickname"`
	Type       string `json:"client_type"`
	Platform   string `json:"client_platform"`
	Version    string `json:"client_version"`
}

func (c *Client) GetClientListWithInfo() ([]ClientList, error) {
	resp, err := c.sendRequest("/clientlist?-info")
	if err != nil {
		return nil, err
	}
	return decodeResponse[[]ClientList](resp)
}
