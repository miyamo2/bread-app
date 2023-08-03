package config

var Conf Config

type Config struct {
	Contentful Contentful `json:"contentful"`
}

type Contentful struct {
	Token string `json:"token"`
}
