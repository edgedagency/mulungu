package model

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/datastore"
	valid "github.com/asaskevich/govalidator"
	"github.com/edgedagency/mulungu/logger"
	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//Model element used to interact with datastore
type Model struct {
	Key          *datastore.Key `json:"key,omitempty" datastore:"__key__,omitempty"`
	ID           int64          `json:"id,omitempty" datastore:"-"`
	TenantID     int64          `json:"tenantID,omitempty" datastore:"tenantID,omitempty"`
	CreatedDate  time.Time      `json:"createdDate,omitempty" datastore:"createdDate,omitempty"`
	ModifiedDate time.Time      `json:"modifiedDate,omitempty" datastore:"modifiedDate,omitempty"`

	Namespace string          `json:"-" datastore:"-"`
	Kind      string          `json:"-" datastore:"-"`
	Context   context.Context `json:"-" datastore:"-"`

	client *datastore.Client
}

//Validate determins if model is valid based on validation tags using https://github.com/asaskevich/govalidator
func (m *Model) Validate(i interface{}) (bool, map[string]string) {
	_, err := valid.ValidateStruct(i)
	if err != nil {
		return false, valid.ErrorsByField(err)
	}
	return true, nil
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
		logger.Errorf(m.Context, "model", "failed to decode model, %s", err.Error())
		return err
	}
	return nil
}

//HydrateWithMap hydrates model
func (m *Model) HydrateWithMap(data map[string]interface{}, i interface{}) error {
	log.Debugf(m.Context, "hydrating %#v with %#v", i, data)

	b, byteConversationErr := util.InterfaceToByte(data)
	if byteConversationErr != nil {
		logger.Errorf(m.Context, "model", "failed to decode model, %s", byteConversationErr.Error())
		return byteConversationErr
	}

	unmarshallErr := json.Unmarshal(b, i)

	if unmarshallErr != nil {
		logger.Errorf(m.Context, "model", "failed to decode model, %s", unmarshallErr.Error())
		return unmarshallErr
	}
	return nil
}

//UnMarshal uses one interface to populate another
func (m *Model) UnMarshal(src interface{}, dest interface{}) error {
	logger.Debugf(m.Context, "model", "UnMarshalling src%#v, dest:%#v", src, dest)
	bytes, err := json.Marshal(src)
	if err != nil {
		logger.Errorf(m.Context, "model", "failed marshal src: %#v error:%s ", src, err.Error())
		return err
	}
	json.Unmarshal(bytes, dest)
	return nil
}

//Save save model
func (m *Model) Save(parent *datastore.Key, i interface{}) (*datastore.Key, error) {
	logger.Debugf(m.Context, "model", "saving %#v", i)
	key, putErr := m.client.Put(m.Context, m.GetKey(parent), i)
	if putErr != nil {
		logger.Errorf(m.Context, "model", "failed to store model, %s", putErr.Error())
		return nil, putErr
	}

	logger.Debugf(m.Context, "model", "saved %#v with key %#v", i, key)

	return key, nil
}

//SetClient instantiates and sets up client on datastore
func (m *Model) SetClient() {
	client, clientErr := datastore.NewClient(m.Context, appengine.AppID(m.Context))
	if clientErr != nil {
		logger.Errorf(m.Context, "model", "failed to create client, %s", clientErr.Error())
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
