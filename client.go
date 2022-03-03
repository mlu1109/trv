package trv

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	url    string
	client *http.Client
}

func NewDefaultClient() *Client {
	return &Client{
		url:    "https://api.trafikinfo.trafikverket.se/v2/data.json",
		client: http.DefaultClient,
	}
}

func NewClient(client *http.Client) *Client {
	return &Client{
		url:    "https://api.trafikinfo.trafikverket.se/v2/data.json",
		client: client,
	}
}

func (c *Client) WithURL(url string) *Client {
	c.url = url
	return c
}

func (c *Client) Post(r *Request) (*Response, error) {
	xml, err := xml.Marshal(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.url, bytes.NewReader(xml))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "text/xml")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return c.unmarshalResponse(resp)
}

func (c *Client) unmarshalResponse(resp *http.Response) (*Response, error) {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ct := strings.ToLower(resp.Header.Get("content-type"))
	var response Response
	if ct == "" || strings.Contains(ct, "json") {
		err = json.Unmarshal(respBody, &response)
		return &response, err
	} else if strings.Contains(ct, "xml") {
		return nil, fmt.Errorf("xml is not suppored") // xml is not supported as the xml package unfortunately does not support marshaling to interface{}
	} else {
		return nil, fmt.Errorf("unrecognized content type: %s", ct)
	}
}
