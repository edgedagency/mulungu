package core

import (
	"net/http"
	"time"

	"google.golang.org/appengine"

	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
)

//DatastoreQuery Datastore query
type DatastoreQuery struct {
	Filter []DatastoreFilter
}

//AddFilter adds filter to query
func (q *DatastoreQuery) AddFilter(datastoreFilter DatastoreFilter) {
	q.Filter = append(q.Filter, datastoreFilter)
}

//DatastoreFilter filter object
type DatastoreFilter struct {
	Key       string `json:"key"`
	Operation string `json:"operation"`
	Value     string `json:"value"`
}

//DatastoreModel struct representing a cloud function
type DatastoreModel struct {
	Namespace          string                 `json:"namespace"`
	Kind               string                 `json:"kind"`
	Record             map[string]interface{} `json:"record"`
	ExcludeFromIndexes []string               `json:"excludeFromIndexes"`
	Query              *DatastoreQuery        `json:"query"`
	Context            context.Context        `json:"-"`
}

//NewDatastoreModel returns a new datastore request struct pointer
func NewDatastoreModel(ctx context.Context, namespace, kind string, record map[string]interface{}, excludeFromIndexes []string) *DatastoreModel {
	datastoreRequest := &DatastoreModel{Context: ctx, Namespace: namespace, Kind: kind, Record: record, ExcludeFromIndexes: excludeFromIndexes}
	return datastoreRequest
}

//JSONBytes returns DatastoreCloudFunction as a transportable json byte
func (ds *DatastoreModel) JSONBytes() []byte {
	bytes, err := util.InterfaceToByte(ds)
	if err != nil {
		return nil
	}
	return bytes
}

//Timestamp timestamps record
func (ds *DatastoreModel) Timestamp() {
	if _, ok := ds.Record["createdDate"]; !ok {
		ds.Record["createdDate"] = time.Now()
	}
	ds.Record["modifiedDate"] = time.Now()
}

//Save timestamps record
func (ds *DatastoreModel) Save() (map[string]interface{}, error) {
	return ds.execute(http.MethodPost, nil)
}

//Update update model by id
func (ds *DatastoreModel) Update(id string) (map[string]interface{}, error) {
	return ds.execute(http.MethodPut, map[string]string{"id": id})
}

//Delete delete model by id
func (ds *DatastoreModel) Delete(id string) (map[string]interface{}, error) {
	return ds.execute(http.MethodDelete, map[string]string{"id": id})
}

//Get execute query on datastore
func (ds *DatastoreModel) Get(searchParams map[string]string) (map[string]interface{}, error) {
	return ds.execute(http.MethodGet, searchParams)
}

func (ds *DatastoreModel) execute(method string, searchParams map[string]string) (map[string]interface{}, error) {

	request, requestErr := util.HTTPNewRequest(ds.Context, method, util.CloudFunctionGetPath("us-central1", appengine.AppID(ds.Context), "dbdatastore"), ds.JSONBytes(), searchParams)
	if requestErr != nil {
		return nil, requestErr
	}

	response, responseErr := util.HTTPRequest(ds.Context, request)
	if responseErr != nil {
		return nil, responseErr
	}

	return util.ResponseToMap(response)
}
