package core

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	valid "github.com/asaskevich/govalidator"
	"github.com/edgedagency/mulungu/logger"
	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

//Model element used to interact with datastore
type Model struct {
	Key          *datastore.Key  `json:"key,omitempty" datastore:"__key__,omitempty"`
	ID           int64           `json:"id,omitempty" datastore:"-"`
	TenantID     int64           `json:"tenantID,omitempty" datastore:"tenantID,omitempty"`
	OwnerID      int64           `json:"ownerID,omitempty" datastore:"ownerID,omitempty"`
	ParentID     int64           `json:"parentID,omitempty" datastore:"parentID,omitempty"`
	Data         Dynamic         `json:"data,omitempty" datastore:"data,omitempty"`
	Status       string          `json:"status" datastore:"status"`
	CreatedDate  time.Time       `json:"createdDate,omitempty" datastore:"createdDate,omitempty"`
	ModifiedDate time.Time       `json:"modifiedDate,omitempty" datastore:"modifiedDate,omitempty"`
	Namespace    string          `json:"-" datastore:"-"`
	Kind         string          `json:"-" datastore:"-"`
	Context      context.Context `json:"-" datastore:"-"`
	client       *datastore.Client
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
func (m *Model) GetKey(id int64, parent *datastore.Key) *datastore.Key {
	key := datastore.IDKey(m.Kind, id, parent)
	key.Namespace = m.Namespace
	return key
}

//Save save model
func (m *Model) Save(parent *datastore.Key, i interface{}) (*datastore.Key, error) {
	logger.Debugf(m.Context, "model", "saving %#v", i)
	key, putErr := m.client.Put(m.Context, m.GetKey(0, parent), i)
	if putErr != nil {
		logger.Errorf(m.Context, "model", "failed to store record, %s", putErr.Error())
		return nil, putErr
	}

	logger.Debugf(m.Context, "model", "saved record %#v with key %#v", i, key)

	return key, nil
}

//Update save model
func (m *Model) Update(id int64, parent *datastore.Key, i interface{}) (*datastore.Key, error) {
	logger.Debugf(m.Context, "model", "updating record with id %d with %#v", id, i)
	key, putErr := m.client.Put(m.Context, m.GetKey(id, parent), i)
	if putErr != nil {
		logger.Errorf(m.Context, "model", "failed to update record, %s", putErr.Error())
		return nil, putErr
	}

	logger.Debugf(m.Context, "model", "updated record %#v with key %#v", i, key)

	return key, nil
}

//DeleteByID deletes user by id
func (m *Model) DeleteByID(id int64, parent *datastore.Key) error {
	m.Identify(datastore.IDKey(m.Kind, id, parent))
	return m.client.Delete(m.Context, m.Key)
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

//Hydrate hydrates model
func (m *Model) Hydrate(readCloser io.ReadCloser, i interface{}) error {
	logger.Debugf(m.Context, "model", "hydrating %#v", i)

	err := json.NewDecoder(readCloser).Decode(i)
	if err != nil {
		logger.Errorf(m.Context, "model", "failed to decode model, %s", err.Error())
		return err
	}
	return nil
}

//HydrateWithMap hydrates model
func (m *Model) HydrateWithMap(data map[string]interface{}, i interface{}) error {

	dataString := util.MapInterfaceToJSONString(data)
	logger.Debugf(m.Context, "model", "hydrating: %#v with: %#v data-string: %#v", i, data, dataString)

	d := json.NewDecoder(strings.NewReader(dataString))
	d.UseNumber()

	if decoderErr := d.Decode(i); decoderErr != nil {
		logger.Errorf(m.Context, "model", "failed to decode, %s", decoderErr.Error())
		return decoderErr
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

//SetClient instantiates and sets up client on datastore
func (m *Model) SetClient() {
	client, clientErr := datastore.NewClient(m.Context, appengine.AppID(m.Context))
	if clientErr != nil {
		logger.Errorf(m.Context, "model", "failed to create client, %s", clientErr.Error())
		panic(fmt.Errorf("failed to create client, %s", clientErr.Error()))
	}
	m.client = client
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

//IndexPut creates a search index
func (m *Model) IndexPut(id int64, kind string, src interface{}) error {
	logger.Debugf(m.Context, "model", "creating search index, kind %s,  id %s", kind, id)
	index, openError := search.Open(kind)

	if openError != nil {
		logger.Errorf(m.Context, "model", "failed to open search index, %s", openError.Error())
		return openError
	}
	_, putError := index.Put(m.Context, strconv.FormatInt(id, 10), src)
	if putError != nil {
		logger.Errorf(m.Context, "model", "failed to create search index %s", putError.Error())
		return putError
	}
	return nil
}
