package model

type Stage struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}
