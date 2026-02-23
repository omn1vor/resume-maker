package model

import "strings"

type Resume struct {
	Name     string `yaml:"name"`
	Position string `yaml:"position"`
	Address  string `yaml:"address"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
	LinkedIn string `yaml:"linkedIn"`
	Github   string `yaml:"github"`
	PhotoURL string `yaml:"photoUrl"`

	Languages []string `yaml:"languages"`
	Skills    []string `yaml:"skills"`

	Experience []Experience `yaml:"experience"`
	Education  []Education  `yaml:"education"`
	Courses    []Course     `yaml:"courses"`
	Projects   []Project    `yaml:"projects"`

	Profile string `yaml:"profile"`
}

type Experience struct {
	Company  string   `yaml:"company"`
	Position string   `yaml:"position"`
	Start    string   `yaml:"start"`
	End      string   `yaml:"end"`
	Skills   []string `yaml:"skills"`
	Comment  string   `yaml:"comment"`
}

type Education struct {
	Year        string `yaml:"year"`
	Institution string `yaml:"institution"`
	Major       string `yaml:"major"`
}

type Course struct {
	Year    string `yaml:"year"`
	Company string `yaml:"company"`
	Name    string `yaml:"name"`
}

type Project struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	URL         string   `yaml:"url"`
	Skills      []string `yaml:"skills"`
}

func (r *Resume) FileName() string {
	name := strings.ReplaceAll(strings.TrimSpace(r.Name), "  ", " ")
	name = strings.ToLower(name)
	lines := strings.Split(name, " ")

	position := strings.ReplaceAll(strings.TrimSpace(r.Position), "  ", " ")
	position = strings.ToLower(position)
	lines = append(lines, strings.Split(position, " ")...)

	return strings.Join(lines, "_") + ".pdf"
}
