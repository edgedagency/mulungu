package model

import (
	"encoding/json"
	"io"
	"time"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//Model element used to interact with datastore
type Model struct {
	Key          *datastore.Key  `json:"key,omitempty" datastore:"__key__,omitempty"`
	ID           int64           `json:"id,omitempty" datastore:"_,omitempty"`
	CreatedDate  time.Time       `json:"createdDate" datastore:"createdDate,omitempty"`
	ModifiedDate time.Time       `json:"modifiedDate" datastore:"modifiedDate,omitempty"`
	Namespace    string          `json:"_" datastore:"_"`
	Kind         string          `json:"_" datastore:"_"`
	Context      context.Context `json:"_" datastore:"_"`
}

//GetKey generates new key
func (m *Model) GetKey(parent *datastore.Key) *datastore.Key {
	key := datastore.IDKey(m.Kind, 0, parent)
	key.Namespace = m.Namespace
	return key
}

//GetKind returns entity kind
func (m *Model) GetKind() string {
	return m.Kind
}

//Hydrate hydrates model
func (m *Model) Hydrate(readCloser io.ReadCloser, i interface{}) error {
	log.Debugf(m.Context, "hydrating %#v", i)

	err := json.NewDecoder(readCloser).Decode(i)
	if err != nil {
		log.Errorf(m.Context, "failed to decode model, %s", err.Error())
		return err
	}
	return nil
}

// //GetModel retuens this model
// func (m *Model) GetModel() *Model {
// 	return m
// }

//Save save model
func (m *Model) Save(parent *datastore.Key, i interface{}) (*datastore.Key, error) {
	client, clientErr := m.GetClient()
	if clientErr != nil {
		log.Errorf(m.Context, "failed to create client, %s", clientErr.Error())
		return nil, clientErr
	}

	//fixme:make this part of model code somehow
	// i.GetModel().CreatedDate = time.Now()
	// i.GetModel().ModifiedDate = time.Now()

	log.Debugf(m.Context, "%#v", i)
	key, putErr := client.Put(m.Context, m.GetKey(parent), i)
	if putErr != nil {
		log.Errorf(m.Context, "failed to store model, %s", putErr.Error())
		return nil, putErr
	}

	return key, nil
}

//GetClient returns client which cna be used to communicate with datastore
func (m *Model) GetClient() (*datastore.Client, error) {
	client, clientErr := datastore.NewClient(m.Context, appengine.AppID(m.Context))
	if clientErr != nil {
		log.Errorf(m.Context, "failed to create client, %s", clientErr.Error())
		return client, clientErr
	}
	return client, nil
}
