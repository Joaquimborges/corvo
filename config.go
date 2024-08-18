package main

type Config struct {
	PostCard          string
	AuthorizationCode string
	UrlMapper         map[urlKey]string
}
