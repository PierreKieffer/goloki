package goloki

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type LogObject struct {
	Streams []StreamObject `json:"streams"`
}

type StreamObject struct {
	Stream map[string]interface{} `json:"stream"`
	Values []Value                `json:"values"`
}

type Value []string

func Log(message string, optLabels ...map[string]interface{}) *LogObject {

	ts := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	var labels = make(map[string]interface{})

	if len(optLabels) > 0 {
		labels = optLabels[0]
	}

	log := LogObject{
		Streams: []StreamObject{
			StreamObject{
				Stream: labels,
				Values: []Value{
					Value{ts, message},
				},
			},
		},
	}
	return &log
}

func LogGroup(messages []string, optLabels ...map[string]interface{}) *LogObject {

	ts := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	var labels = make(map[string]interface{})

	var values []Value
	for _, m := range messages {
		values = append(values, Value{ts, m})
	}

	if len(optLabels) > 0 {
		labels = optLabels[0]
	}

	log := LogObject{
		Streams: []StreamObject{
			StreamObject{
				Stream: labels,
				Values: values,
			},
		},
	}
	return &log
}

func (l *LogObject) Push(lokiUrl string) error {
	/*
	   Send log entries to Loki
	*/

	endpoint := fmt.Sprintf("%v/loki/api/v1/push", lokiUrl)

	err := PostRequest(endpoint, l)
	if err != nil {
		return err
	}
	return nil
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

	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("middleWare.PostRequest http status code %v", resp.StatusCode))
	}
	return nil
}
