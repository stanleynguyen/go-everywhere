package lighthttpcli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	StateOn  = "ON"
	StateOff = "OFF"
)

type LightHttpCli struct {
	url        string
	httpClient *http.Client
}

func NewCli(serverURL string) LightHttpCli {
	return LightHttpCli{
		url:        serverURL,
		httpClient: http.DefaultClient,
	}
}

func (c LightHttpCli) WithHttpClient(cli *http.Client) LightHttpCli {
	c.httpClient = cli
	return c
}

func (c LightHttpCli) GetState() (string, error) {
	endpoint := fmt.Sprintf("%s/led", c.url)
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return "OFF", err
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "OFF", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "OFF", errors.New(string(respByte))
	}

	return string(respByte), nil
}

func (c LightHttpCli) SetState(state string) error {
	endpoint := fmt.Sprintf("%s/%s", c.url, strings.ToLower(state))
	req, _ := http.NewRequest(http.MethodPost, endpoint, nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(respByte))
	}

	return nil
}
