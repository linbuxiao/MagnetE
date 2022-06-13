package model

type RuleRepo struct {
	Name   string `json:"name" yaml:"name"`
	Source string `json:"source" yaml:"source"`
	Item   string `json:"item" yaml:"item"`
	Title  string `json:"title" yaml:"title"`
	Magnet string `json:"magnet" yaml:"magnet"`
	Size   string `json:"size" yaml:"size"`
}
