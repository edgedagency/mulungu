package core

import "github.com/edgedagency/mulungu/util"

//Dataset a record database record
type Dataset map[string]interface{}

//NewDataset returns a record
func NewDataset(data map[string]interface{}) *Dataset {
	record := make(Dataset)
	for k, v := range data {
		record[k] = v
	}
	return &record
}

func (r *Dataset) Get(key string) interface{} {
	if val, ok := (*r)[key]; ok {
		return val
	}
	return nil
}

func (r *Dataset) GetString(key string) string {
	if val, ok := (*r)[key]; ok {
		return util.ToString(val)
	}
	return ""
}

func (r *Dataset) GetInt64(key string) int64 {
	if val, ok := (*r)[key]; ok {
		return util.NumberizeString(val).(int64)
	}
	return 0
}

func (r *Dataset) GetInt32(key string) int {
	if val, ok := (*r)[key]; ok {
		return util.NumberizeString(val).(int)
	}
	return 0
}
