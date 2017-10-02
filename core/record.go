package core

import "github.com/edgedagency/mulungu/util"

//Record a record database record
type Record map[string]interface{}

//NewRecord returns a record
func NewRecord(data map[string]interface{}) *Record {
	record := make(Record)
	for k, v := range data {
		record[k] = v
	}
	return &record
}

func (r *Record) Get(key string) interface{} {
	if val, ok := (*r)[key]; ok {
		return val
	}
	return nil
}

func (r *Record) GetString(key string) string {
	if val, ok := (*r)[key]; ok {
		return util.ToString(val)
	}
	return ""
}

func (r *Record) GetInt64(key string) int64 {
	if val, ok := (*r)[key]; ok {
		return util.NumberizeString(val).(int64)
	}
	return 0
}

func (r *Record) GetInt32(key string) int {
	if val, ok := (*r)[key]; ok {
		return util.NumberizeString(val).(int)
	}
	return 0
}
