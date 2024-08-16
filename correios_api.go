package main

type WebServices interface {
}

type webServices struct {
	client *restClient
	config *Config
}

func NewCorreiosWebServices(config *Config) WebServices {
	return &webServices{
		client: &restClient{},
		config: config,
	}
}
