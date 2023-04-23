package cache

import (
	"errors"
	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/types"
)

var lastPayload *types.LiveData

var ErrNoLastPayload = errors.New("cache: no last payload available")

func LastPayload() (*types.LiveData, error) {
	if lastPayload == nil {
		return nil, ErrNoLastPayload
	}
	return lastPayload, nil
}

func SetLastPayload(payload *types.LiveData) {
	lastPayload = payload
}
