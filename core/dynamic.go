package core

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/edgedagency/mulungu/util"
)

//Dynamic represents a user contacts
type Dynamic map[string]interface{}

//Load function from PropertyLoaderInterface helps datastore load this object
func (d *Dynamic) Load(dp []datastore.Property) error {
	*d = make(Dynamic)
	for _, property := range dp {
		(*d)[property.Name] = property.Value
	}
	return nil
}

//Save function from PropertyLoaderInterface helps datastore save this object
func (d *Dynamic) Save() ([]datastore.Property, error) {
	propertise := []datastore.Property{}
	for name, value := range *d {
		fmt.Println(fmt.Printf("property name: %s value: %v", name, value))
		propertise = d.AppendProperty(propertise, name, true, value)
	}
	return propertise, nil
}

//AppendProperty converts data to interface
func (d *Dynamic) AppendProperty(propertise []datastore.Property, name string, index bool, value interface{}) []datastore.Property {
	if util.IsDatastoreAcceptableType(value) {
		return append(propertise, util.GetDatastoreProperty(name, index, value))
	}
	return propertise
}
