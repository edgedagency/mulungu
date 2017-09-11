package util

//PubSubData obtain data item by key
func PubSubData(data map[string]interface{}, messageKey, key string) interface{} {
	if dataElement, ok := data[messageKey]; ok {
		if val, ok := dataElement.(map[string]interface{})[key]; ok {
			return val
		}
	}
	return nil
}

// func PubSubSubscription(subscription string) map[string]string {
// 	subscription = subscription[strings.LastIndex(subscription, "/"):len(subscription)]
//
// }
