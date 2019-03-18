package main

import "github.com/cortexproject/cortex/pkg/util/flagext"

import (
	"fmt"
	"strconv"
	"time"
)

type lokiConfig struct {
	url       flagext.URLValue
	batchWait time.Duration
	batchSize int
}

func getLokiConfig(url string, batchWait string, batchSize string) (*lokiConfig, error) {
	lc := &lokiConfig{}
	var clientURL flagext.URLValue
	if url == "" {
		url = "http://localhost:3100/api/prom/push"
	}
	err := clientURL.Set(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse client URL")
	}
	lc.url = clientURL

	batchWaitValue, err := strconv.Atoi(batchWait)
	if err != nil || batchWait == "" {
		batchWaitValue = 10
	}
	lc.batchWait = time.Duration(batchWaitValue) * time.Millisecond

	batchSizeValue, err := strconv.Atoi(batchSize)
	if err != nil || batchSize == "" {
		batchSizeValue = 10
	}
	lc.batchSize = batchSizeValue * 1024

	return lc, nil
}
