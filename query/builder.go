package query

import (
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//query constants
const (
	NoSorting   = ""
	NoFiltering = ""
)

//Builder struct holding query building logic
type Builder struct {
	Context context.Context
	Query   *datastore.Query
}

//NewQueryBuilder returns a new query.Builder
func NewQueryBuilder(ctx context.Context) *Builder {
	builder := &Builder{Context: ctx}
	return builder
}

//Build builds a query based on provided parameters
func (b *Builder) Build(kind, filter, sort string) {
	b.Query = datastore.NewQuery(kind)
	b.Filter(filter)
}

//Filter adds and processes filtering logic
func (b *Builder) Filter(filter string) {
	if filter != "" {
		log.Debugf(b.Context, "query filter %s", filter)
		filters := strings.Split(filter, ",")
		log.Debugf(b.Context, "filters parts %#v", filters)
		for _, filterPart := range filters {
			log.Debugf(b.Context, "filterPart: %s", filterPart)
			queryParts := strings.Split(filterPart, ":")
			log.Debugf(b.Context, "filterParts: filter: %s value:%s", queryParts[0], queryParts[1])
			b.Query = b.Query.Filter(queryParts[0], util.NumberizeString(queryParts[1]))
		}
	}

	log.Debugf(b.Context, "query after filters: %#v", b.Query)
}

//Order addds order by clause
func (b *Builder) Order(order string) {
	b.Query = b.Query.Order(order)
}
