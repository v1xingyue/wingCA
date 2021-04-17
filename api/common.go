package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func downloadFileHandler(fileName, savedName string) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if data, err := ioutil.ReadFile(fileName); err == nil {
			w.Header().Add("Content-Type", "application/force-download")
			w.Header().Add("Content-Disposition", "attachment;filename="+savedName)
			w.Header().Add("Content-Transfer-Encoding", "binary")
			w.Write(data)
		} else {
			w.Write([]byte(fmt.Sprintf("%s", err)))
		}
	}

}
