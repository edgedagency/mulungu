package provider

//DataProvider interface for data providers
type DataProvider interface {
	//Save save record
	Save(collectionName string, data []byte) (map[string]interface{}, error)

	//Update update record
	Update(collectionName, id string, data []byte) (map[string]interface{}, error)

	//Delete delete record
	Delete(collectionName, id string) (map[string]interface{}, error)

	//Find find record based on provided identifier
	Find(collectionName, id string) (map[string]interface{}, error)

	//FindAll search for and obtain records
	FindAll(collectionName, filter, sort, order string, limit int, page int, selects []string) ([]interface{}, error)
}
