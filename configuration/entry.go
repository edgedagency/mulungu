package configuration

import (
	"golang.org/x/net/context"

	"github.com/edgedagency/mulungu/core"
)

//Entry this is a representation of a configuration entry
//
//e.g. email.from.address = mince@example.com
//e.g. application.name = Mince Inc Super App
type Entry struct {
	core.Model

	Key   string `json:"key" datastore:"key"`
	Value string `json:"value" datastore:"value"`
}

//NewConfigurationEntryModel instantiates a new user model
func NewConfigurationEntryModel(context context.Context, namespace string) *Entry {
	m := &Entry{}
	m.Init(context, "ConfigurationEntry", namespace)
	return m
}

//Get retireves a configuration
func (e *Entry) Get(key string) string {
	//get configuration entry by key
	return ""
}

//Set sets or updates a configuration
func (e *Entry) Set(key string, value string) *Entry {

	//set  configuration entry with key
	// overrides existing entry, therefore check ig an entry with key exists update if true create new if false
	return nil
}
