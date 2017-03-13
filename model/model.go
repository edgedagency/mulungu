package model

import (
	"encoding/json"
	"fmt"
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
	ID           int64           `json:"id,omitempty" datastore:"-"`
	CreatedDate  time.Time       `json:"createdDate" datastore:"createdDate,omitempty"`
	ModifiedDate time.Time       `json:"modifiedDate" datastore:"modifiedDate,omitempty"`
	Namespace    string          `json:"-" datastore:"-"`
	Kind         string          `json:"-" datastore:"-"`
	Context      context.Context `json:"-" datastore:"-"`

	client *datastore.Client
}

//Init initializes the model with necessary logic
func (m *Model) Init(context context.Context, kind string, namespace string) {
	m.Context = context
	m.Kind = kind
	m.Namespace = namespace
	m.SetClient()
}

//Client returns instentiated client
func (m *Model) Client() *datastore.Client {
	return m.client
}

//GetKey generates new key
func (m *Model) GetKey(parent *datastore.Key) *datastore.Key {
	key := datastore.IDKey(m.Kind, 0, parent)
	key.Namespace = m.Namespace
	return key
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

//UnMarshal uses one interface to populate another
func (m *Model) UnMarshal(src interface{}, dest interface{}) error {
	log.Debugf(m.Context, "UnMarshalling src%#v, dest:%#v", src, dest)
	bytes, err := json.Marshal(src)
	if err != nil {
		log.Errorf(m.Context, "failed marshal src: %#v error:%s ", src, err.Error())
		return err
	}
	json.Unmarshal(bytes, dest)
	return nil
}

//Save save model
func (m *Model) Save(parent *datastore.Key, i interface{}) (*datastore.Key, error) {
	log.Debugf(m.Context, "%#v", i)
	key, putErr := m.client.Put(m.Context, m.GetKey(parent), i)
	if putErr != nil {
		log.Errorf(m.Context, "failed to store model, %s", putErr.Error())
		return nil, putErr
	}

	return key, nil
}

//SetClient instantiates and sets up client on datastore
func (m *Model) SetClient() {
	client, clientErr := datastore.NewClient(m.Context, appengine.AppID(m.Context))
	if clientErr != nil {
		log.Errorf(m.Context, "failed to create client, %s", clientErr.Error())
		panic(fmt.Errorf("failed to create client, %s", clientErr.Error()))
	}
	m.client = client
}

//Run runs passwd query and returns results
func (m *Model) Run(query *datastore.Query) *datastore.Iterator {
	log.Debugf(m.Context, "running query, %#v", query)
	return m.client.Run(m.Context, query)
}

//FindByID finds model my id
func (m *Model) FindByID(id int64, parent *datastore.Key, destination interface{}) error {
	m.Identify(datastore.IDKey(m.Kind, id, parent))
	return m.client.Get(m.Context, m.Key, destination)

}

//DeleteByID deletes user by id
func (m *Model) DeleteByID(id int64, parent *datastore.Key) error {
	m.Identify(datastore.IDKey(m.Kind, id, parent))
	return m.client.Delete(m.Context, m.Key)
}

//Identify attachs model identifications key and id, need on retrival of information
func (m *Model) Identify(key *datastore.Key) {
	m.ID = key.ID
	if m.Key == nil {
		m.Key = key
		m.Key.Namespace = m.Namespace
	}
}

//Timestamp timestamps model
func (m *Model) Timestamp() {
	if m.CreatedDate.IsZero() {
		m.CreatedDate = time.Now()
	}
	m.ModifiedDate = time.Now()
}

//IsNil checks to see if this model is empty, since model can't be compared to nil
func (m *Model) IsNil() bool {
	return m.CreatedDate.IsZero() == true && m.ModifiedDate.IsZero() == true && m.Key == nil
}
