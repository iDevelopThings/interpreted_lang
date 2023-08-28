package utilities

import (
	"github.com/goccy/go-json"

	"arc/log"
)

func ToJson(data any) []byte {
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("cannot marshal data: %v", err)
	}
	return d
}
