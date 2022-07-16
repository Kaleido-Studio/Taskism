package models

import "github.com/kamva/mgm/v3"

type Task struct {
	mgm.DefaultModel `bson:",inline"`

	Name    string `json:"name" bson:"name"`
	Content string `json:"content" bson:"content"`
}

type Project struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string          `json:"name" bson:"name"`
	Group            string          `json:"group" bson:"group"`
	Description      string          `json:"description" bson:"description"`
	Tags             []string        `json:"tags" bson:"tags"`
	Lanes            map[string]Task `json:"lanes" bson:"lanes"`
}

func NewProject(name string) *Project {
	return &Project{
		Name: name,
	}
}
