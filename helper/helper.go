package helper

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Marshaller post body into type while checking for errs
func Marshaller(r *http.Request, s interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, s); err != nil {
		return err
	}
	return nil
}
