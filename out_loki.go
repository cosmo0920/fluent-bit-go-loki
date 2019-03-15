package main

import "github.com/fluent/fluent-bit-go/output"
import "github.com/grafana/loki/pkg/promtail/client"
import "github.com/sirupsen/logrus"
import kit "github.com/go-kit/kit/log/logrus"
import "github.com/prometheus/common/model"
import "github.com/cortexproject/cortex/pkg/util/flagext"

import (
	"C"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"unsafe"
)

var loki *client.Client
var ls model.LabelSet

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	return output.FLBPluginRegister(ctx, "loki", "Loki GO!")
}

//export FLBPluginInit
// (fluentbit will call this)
// ctx (context) pointer to fluentbit context (state/ c code)
func FLBPluginInit(ctx unsafe.Pointer) int {
	// Example to retrieve an optional configuration parameter
	url := output.FLBPluginConfigKey(ctx, "url")
	var clientURL flagext.URLValue
	err := clientURL.Set(url)
	if err != nil {
		log.Fatalf("Failed to parse client URL")
	}
	fmt.Printf("[flb-go] plugin URL parameter = '%s'\n", url)

	cfg := client.Config{}
	// Init everything with default values.
	flagext.RegisterFlags(&cfg)

	// Override some of those defaults
	cfg.URL = clientURL
	cfg.BatchWait = 10 * time.Millisecond
	cfg.BatchSize = 10 * 1024

	log := logrus.New()

	loki, err = client.New(cfg, kit.NewLogrusLogger(log))
	if err != nil {
		log.Fatalf("client.New: %s\n", err)
	}
	ls = model.LabelSet{"job": "fluent-bit"}

	return output.FLB_OK
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	var ret int
	var ts interface{}
	var record map[interface{}]interface{}

	dec := output.NewDecoder(data, int(length))

	for {
		ret, ts, record = output.GetRecord(dec)
		if ret != 0 {
			break
		}

		// Get timestamp
		timestamp := ts.(output.FLBTime).Time

		js, err := createJSON(timestamp, record)
		if err != nil {
			fmt.Errorf("error creating message for Grafana Loki: %v", err)
			continue
		}

		err = loki.Handle(ls, timestamp, string(js))
		if err != nil {
			fmt.Errorf("error sending message for Grafana Loki: %v", err)
			return output.FLB_RETRY
		}
	}

	// Return options:
	//
	// output.FLB_OK    = data have been processed.
	// output.FLB_ERROR = unrecoverable error, do not try this again.
	// output.FLB_RETRY = retry to flush later.
	return output.FLB_OK
}

func createJSON(timestamp time.Time, record map[interface{}]interface{}) (string, error) {
	m := make(map[string]interface{})

	for k, v := range record {
		switch t := v.(type) {
		case []byte:
			// prevent encoding to base64
			m[k.(string)] = string(t)
		default:
			m[k.(string)] = v
		}
	}

	js, err := json.Marshal(m)
	if err != nil {
		return "{}", err
	}

	return string(js), nil
}

//export FLBPluginExit
func FLBPluginExit() int {
	loki.Stop()
	return output.FLB_OK
}

func main() {
}
