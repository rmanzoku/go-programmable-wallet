package programmablewallet

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	V1Endpoint = "https://api.circle.com/v1"
)

type ProgrammableWallet struct {
	Endpoint string
	BasePath string
	Debug    bool

	apikey string
}

func NewProgrammableWallet(apikey string) (*ProgrammableWallet, error) {
	p := &ProgrammableWallet{
		Endpoint: V1Endpoint,
		BasePath: "/w3s/",
		Debug:    false,
		apikey:   apikey,
	}
	return p, nil
}

func (p *ProgrammableWallet) EnableDebug() {
	p.Debug = true
}

// Get configuration for entity
func (p *ProgrammableWallet) GetConfigurationForEntity(ctx context.Context) (string, error) {
	path := "config/entity"

	data, err := p.Get(ctx, path)
	if err != nil {
		return "", err
	}

	d := struct {
		AppId string `json:"appId"`
	}{}
	return d.AppId, json.Unmarshal(data, &d)
}

type Response struct {
	Response struct {
		Status struct {
			Version string `json:"version"`
			Message string `json:"message"`
		} `json:"status"`
		Data json.RawMessage `json:"data"`
	} `json:"response"`
}

func (p *ProgrammableWallet) Get(ctx context.Context, path string) (json.RawMessage, error) {

	url := p.Endpoint + p.BasePath + path

	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+p.apikey)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r.Response.Data, nil
}
