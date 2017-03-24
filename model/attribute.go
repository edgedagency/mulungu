package model

import "github.com/edgedagency/mulungu/constant"

//Attribute can be used to present attributes which can be attached to models
type Attribute struct {
	Model
	Name        string          `json:"name" datastore:"name"`
	Value       string          `json:"value" datastore:"value,noindex"`
	Description string          `json:"description,omitempty" datastore:"description"`
	Type        string          `json:"type" datastore:"type"` //What type is this attributes (number, string, float, list, matrix)
	Attributes  []Attribute     `json:"attributes,omitempty" datastore:"attributes"`
	Status      constant.Status `json:"status" datastore:"status"`
}
