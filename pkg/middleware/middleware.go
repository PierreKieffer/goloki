package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	// "goloki/pkg/configuration"
	"net/http"
)

type Log struct {
	Streams []StreamObject `json:"streams"`
}
type StreamObject struct {
	Entries []EntryObject          `json:"entries"`
	Labels  map[string]interface{} `json:"labels"`
}
type EntryObject struct {
	Line      string `json:"line"`
	Timestamp string `json:"ts"`
}

/*
{
    "streams": [
        {
            "entries": [
                {
                    "line": "fizzbuzz",
                    "ts": "2021-06-08T05:28:06.801064-04:00"
                }
            ],
            "labels": "{foo=\"bar\"}"
        }
    ]
}


curl -H "Content-Type: application/json" -XPOST -s "http://23.251.135.162:3100/loki/api/v1/push" --data-raw \
  '{"streams": [{ "labels": "{foo=\"bar\"}", "entries": [{ "ts": "2021-06-08T05:28:06.801064-04:00", "line": "fizzbuzz" }] }]}'
*/

// func (l *Log) Push(config *configuration.Configuration) {
// /*
// Send log entries to Loki
// */

// endpoint := fmt.Sprintf("%v/loki/api/v1/push", config.LokiUrl)

// }

func InitLog(message string, labels map[string]interface{}) *Log {
	ts := time.Now().UTC().Format(time.RFC3339)
	log := Log{
		Streams: []StreamObject{
			StreamObject{
				Entries: []EntryObject{
					EntryObject{
						Line:      message,
						Timestamp: ts,
					},
				},
				Labels: labels,
			},
		},
	}
	return &log

}

func PostRequest(endpoint string, payload interface{}) error {
	jsonPayload, _ := json.Marshal(payload)
	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("middleWare.PostRequest http status code %v", resp.StatusCode))
	}
	return nil
}
