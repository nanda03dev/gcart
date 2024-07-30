package common

import "github.com/nanda03dev/gnosql_client"

func GetStringValue(document gnosql_client.Document, key string) string {
	value, _ := document[key]
	return value.(string)
}
func GetIntegerValue(document gnosql_client.Document, key string) int {
	value, _ := document[key]
	return value.(int)
}
func GetBoolValue(document gnosql_client.Document, key string) bool {
	value, _ := document[key]
	return value.(bool)
}
func GetValue[T any](document gnosql_client.Document, key string) T {
	value, _ := document[key]
	return value.(T)
}
