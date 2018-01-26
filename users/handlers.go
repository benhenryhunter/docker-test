package adsets

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/davecgh/go-spew/spew"
)

func Render(w http.ResponseWriter, r *http.Request) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting path: " + err.Error()); err != nil {
			panic(err)
		}
	}
	vars := mux.Vars(r)
	query := vars["who"]
	spew.Dump(r.Cookies)
	spew.Dump(r.Header)
	d1 := []byte("<html><p style='color:red'>hello " + query + "</p></html>")
	err = ioutil.WriteFile(dir+"/adsets/html/dat1", d1, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error generating html: " + err.Error()); err != nil {
			panic(err)
		}
	}
	write, err := ioutil.ReadFile(dir + "/adsets/html/dat1")
	w.WriteHeader(http.StatusOK)
	w.Write(write)
}
