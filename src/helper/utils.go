package helper

import (
	"encoding/json"
)

func ResponseToStruct(body []byte, obj interface{}) error {
	if err := json.Unmarshal(body, &obj); err != nil {
		return err
	}

	return nil
}
