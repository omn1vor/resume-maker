package registry

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/omn1vor/resume-maker/internal/model"
)

const (
	templatesDir     = "templates"
	templateFileName = "template.html"
	extHtml          = ".html"
)

type Registry struct {
	templates map[string]model.Template
}

func New() (*Registry, error) {
	r := &Registry{
		templates: map[string]model.Template{},
	}
	err := r.init()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Registry) GetTemplate(name string) (*model.Template, error) {
	temp, ok := r.templates[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}
	return &temp, nil
}

func (r *Registry) List() string {
	list := make([]string, 0, len(r.templates))
	for name := range r.templates {
		list = append(list, fmt.Sprintf("- %s", name))
	}
	slices.Sort(list)
	return strings.Join(list, "\n")
}

func (r *Registry) init() error {
	dirs, err := os.ReadDir(templatesDir)
	if err != nil {
		return fmt.Errorf("failed to read templates dir: %w", err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		name := dir.Name()
		files, err := os.ReadDir(path.Join(templatesDir, dir.Name()))
		if err != nil {
			return fmt.Errorf("failed to read template dir for %s: %w", name, err)
		}
		var found bool
		for _, file := range files {
			info, err := file.Info()
			if err != nil {
				return fmt.Errorf("failed to get info for file %s: %w", file.Name(), err)
			}
			if path.Ext(info.Name()) == extHtml {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("no html file found for template %s", name)
		}
		r.templates[name] = model.Template{
			Name: name,
			Path: path.Join(templatesDir, name, templateFileName),
		}
	}
	return nil
}
