package teamspeak

import "fmt"

type ServerInfo struct {
	Name     string `json:"virtualserver_name"`
	Platform string `json:"virtualserver_platform"`
	Version  string `json:"virtualserver_version"`
}

func (c *Client) GetServerInfo() (ServerInfo, error) {
	resp, err := c.sendRequest("/serverinfo")
	if err != nil {
		return ServerInfo{}, err
	}

	servers, err := decodeResponse[[]ServerInfo](resp)
	if err != nil {
		return ServerInfo{}, err
	}

	if len(servers) == 0 {
		return ServerInfo{}, fmt.Errorf("no server info in response")
	}

	return servers[0], nil
}
