package config

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

type Builder struct {
	Mirror string
	Images []Image
	Image  Image
}

type Image struct {
	Release  string
	Target   string
	Profile  string
	Packages []string
	Files    []string
}

func FromBytes(b []byte) (*Builder, error) {
	cfg := &Builder{}
	err := toml.Unmarshal(b, cfg)
	return cfg, err
}

func (c Builder) ToJSON() (b []byte) {
	b, err := json.Marshal(c)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal json"))
	}
	return
}
