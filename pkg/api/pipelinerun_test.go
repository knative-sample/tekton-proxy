package api

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"

	"encoding/json"
	"io/ioutil"
)

func TestAddPingPipelineRun(t *testing.T) {
	args := AddPipelineRunArgs{
		Name:      fmt.Sprintf("test-%d", time.Now().Unix()),
		Namespace: "tutorial",
	}

	bts, _ := json.Marshal(args)
	url := "http://127.0.0.1:9090/v1/ping/pipelinerun"
	req, err := http.NewRequest("POST", url, bytes.NewReader(bts))
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to read response: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read responseBody error:%s", err.Error())
	}

	t.Logf("response:%s ", bodyBytes)

}
