package drudapi

import "encoding/json"

// Client ...
type Client struct {
	Created string `json:"_created,omitempty"`
	Etag    string `json:"_etag,omitempty"`
	ID      string `json:"_id,omitempty"`
	Updated string `json:"_updated,omitempty"`
	Email   string `json:"email"`
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone"`
}

// Path ...
func (c Client) Path(method string) string {
	var path string

	if method == "POST" {
		path = "client"
	} else {
		path = "client/" + c.Name
	}
	return path
}

// Unmarshal ...
func (c *Client) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, &c)
	return err
}

// JSON ...
func (c Client) JSON() []byte {
	c.ID = ""
	c.Etag = ""
	c.Created = ""
	c.Updated = ""

	jbytes, _ := json.Marshal(c)
	return jbytes
}

// PatchJSON ...
func (c Client) PatchJSON() []byte {
	c.ID = ""
	c.Etag = ""
	c.Created = ""
	c.Updated = ""
	// removing name because it has been setup as the id param in drudapi and cannot be  patched
	c.Name = ""

	jbytes, _ := json.Marshal(c)
	return jbytes
}

// ETAG ...
func (c Client) ETAG() string {
	return c.Etag
}

// ClientList ...
type ClientList struct {
	Items []Client `json:"_items"`
	Meta  struct {
		MaxResults int `json:"max_results"`
		Page       int `json:"page"`
		Total      int `json:"total"`
	} `json:"_meta"`
}

// Path ...
func (c ClientList) Path(method string) string {
	return "client"
}

// Unmarshal ...
func (c *ClientList) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, &c)
	return err
}

// JSON ...
func (c ClientList) JSON() []byte {
	jbytes, _ := json.Marshal(c)
	return jbytes
}

// PatchJSON ...
func (c ClientList) PatchJSON() []byte {
	return []byte("")
}

// ETAG ...
func (c ClientList) ETAG() string {
	return ""
}
