package str

import (
	"encoding/json"
	"strings"
)

func Ptr(s string) *string {
	return &s
}

func StructToJson(o interface{}) string {
	b, _ := json.Marshal(o)
	return string(b)
}

func StructToJsonReader(o interface{}) *strings.Reader {
	return strings.NewReader(StructToJson(o))
}
