package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func ParseBody(r *http.Request, data interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), &data); err == nil {
			return
		}
	}
}