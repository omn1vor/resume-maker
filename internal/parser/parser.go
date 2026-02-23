package parser

import (
	"os"

	"github.com/go-yaml/yaml"
	"github.com/omn1vor/resume-maker/internal/model"
)

func LoadResume(path string) (*model.Resume, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var r model.Resume
	if err := yaml.Unmarshal(data, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
