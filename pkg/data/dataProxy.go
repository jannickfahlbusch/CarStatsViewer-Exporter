package data

import (
	"encoding/json"
	"log"
	"net/http"
	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/data/cache"
)

type Proxyhandler struct {
}

func (handler *Proxyhandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	lastPayload, err := cache.LastPayload()
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusNoContent)
		return
		return
	}

	err = json.NewEncoder(rw).Encode(lastPayload)
	if err != nil {
		log.Println(err)
	}
}
