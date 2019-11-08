package broker

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func truePtr() *bool {
	b := true
	return &b
}

func falsePtr() *bool {
	b := false
	return &b
}

func (b *SpinnakerBroker) ValidateBrokerAPIVersion(version string) error {
	return nil
}

func CreatePipeline(restEndpoint string, pipeline *pipeline) {
	requestBody, _ := json.Marshal(pipeline)
	resp, err := http.Post(restEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func DeletePipeline(restEndpoint string, pipeline *requestBodyDelete) {
	requestBody, _ := json.Marshal(pipeline)
	req, err := http.NewRequest("DELETE", restEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
