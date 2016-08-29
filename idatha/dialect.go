package idatha

import "fmt"

//Dialect used to store and create URL endpoints
type Dialect struct {
}

// Find returns find URL endpoint
func (d *Dialect) Find(database, collection, key string) string {
	dialect := fmt.Sprintf("/_db/%s/_api/document/%s/%s", database, collection, key)
	fmt.Printf("dialect: %s database: %s collection: %s", dialect, database, collection)
	return dialect
}

// Create returns create URL endpoint
func (d *Dialect) Create(database, collection string) string {
	dialect := fmt.Sprintf("/_db/%s/_api/document?collection=%s", database, collection)
	fmt.Printf("dialect: %s database: %s collection: %s", dialect, database, collection)
	return dialect
}

// NewDialect returns create URL endpoint
func (d *Dialect) NewDialect() *Dialect {
	return &Dialect{}
}

// Cursor executes abitary query
func (d *Dialect) Cursor(database string) string {
	return fmt.Sprintf("/_db/%s/_api/cursor", database)
}

// Version returns endpoint to obtain version
func (d *Dialect) Version() string {
	return "/_api/version?details=true"
}
