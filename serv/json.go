package serv

import (
	"encoding/json"
)

// UnmarshalAndCheckValue Warning: will be modified in future
func UnmarshalAndCheckValue(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
