package urlservice

import "context"

type hashResponse struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type getResponse struct {
	Url string `json:"url"`
}

type Url interface {
	HashUrl(context.Context, string) (*hashResponse, error)
	GetUrl(context.Context, string) (*getResponse, error)
}
