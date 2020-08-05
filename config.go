package main

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

type Database struct {
	User string
	Pass string
	Address string
	Port string
	Name string
}

func (d Database) DSN() string {
	return "postgres://" + d.User + ":" + d.Pass + "@" + d.Address + ":" + d.Port + "/" + d.Name + "?sslmode=disable"
}

type Server struct {
	Address string
	Port string
}

func (s Server) Bind() string {
	return s.Address + ":" + s.Port
}

type Config struct {
	Database Database
	Server   Server
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	confContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read config %v", path)
	}

	if err := toml.Unmarshal(confContent, config); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal config")
	}

	return config, err
}
