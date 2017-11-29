package mulungu

import (
	"golang.org/x/net/context"

	"github.com/edgedagency/mulungu/core"
	"github.com/edgedagency/mulungu/logger"
	"github.com/edgedagency/mulungu/util"
)

//Service represenetation of a service
type Service struct {
	context   context.Context
	namespace string
	kind      string
}

//Init initiates this services
func (s *Service) Init(context context.Context, namespace, kind string) {
	s.context = context
	s.namespace = namespace
	s.kind = kind
}

//Kind returns kind
func (s *Service) Kind() string {
	return s.kind
}

//Namespace returns namespace
func (s *Service) Namespace() string {
	return s.namespace
}

//Context returns context
func (s *Service) Context() context.Context {
	return s.context
}

//Find returns application based on id
func (s *Service) Find(id string) (map[string]interface{}, error) {

	//2. save record
	datastoreModel := core.NewDatastoreModel(s.Context(), s.Namespace(), s.Kind(), nil)
	responseRecord, responseError := datastoreModel.Find(id)

	if responseError != nil {
		return nil, responseError
	}

	return responseRecord, nil
}

//FindAll finds all application from datastore
func (s *Service) FindAll(filter string) (interface{}, error) {

	datastoreModel := core.NewDatastoreModel(s.Context(), s.Namespace(), s.Kind(), nil)
	responseRecord, responseError := datastoreModel.FindAll(filter, "", "", 0, 0, nil)

	if responseError != nil {
		return nil, responseError
	}

	return responseRecord, nil
}

//Publish publish to topic
func (s *Service) Publish(topic string, record map[string]interface{}, attributes map[string]string) {
	publishID, publishErr := util.PubSubTopicPublish(s.Context(), util.PubSubTopicID(s.Namespace(), topic), record, attributes)
	if publishErr != nil {
		logger.Errorf(s.Context(), "application service", "failed to publish record created %s", publishErr.Error())
	}
	logger.Debugf(s.Context(), "application service", "published record created, publish id %s", publishID)
}
