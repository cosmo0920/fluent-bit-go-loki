package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/cortexproject/cortex/pkg/util/flagext"
	"github.com/prometheus/common/model"
)

type lokiConfig struct {
	url       flagext.URLValue
	batchWait time.Duration
	batchSize int
	labelSet  model.LabelSet
}

type labelSetJSON struct {
	Labels []struct {
		Key   string `json:"key"`
		Label string `json:"label"`
	} `json:"labels"`
}

func getLokiConfig(url string, batchWait string, batchSize string, labels string) (*lokiConfig, error) {
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
		batchWaitValue = 1
	}
	lc.batchWait = time.Duration(batchWaitValue) * time.Second

	batchSizeValue, err := strconv.Atoi(batchSize)
	if err != nil || batchSize == "" {
		batchSizeValue = 100 * 1024
	}
	lc.batchSize = batchSizeValue

	var labelValues labelSetJSON
	if labels == "" {
		labels = `
{"labels": [{"key": "job", "label": "fluent-bit"}]}
`
	}

	err = json.Unmarshal(([]byte)(labels), &labelValues)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Labels")
	}
	labelSet := make(model.LabelSet)
	for _, v := range labelValues.Labels {
		labelSet[model.LabelName(v.Key)] = model.LabelValue(v.Label)
	}
	lc.labelSet = labelSet

	return lc, nil
}
