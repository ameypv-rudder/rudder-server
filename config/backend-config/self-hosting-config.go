package backendconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rudderlabs/rudder-server/utils/logger"
)

type SelfHosting struct {
}

func (hosting *SelfHosting) SetUp() {
}

func (hosting *SelfHosting) GetWorkspaceIDForWriteKey(writeKey string) string {
	return ""
}

func (hosting *SelfHosting) GetBackendConfig() (SourcesT, bool) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/workspace-config?workspaceToken=%s", configBackendURL, configBackendToken)
	resp, err := client.Get(url)

	var respBody []byte
	if resp != nil && resp.Body != nil {
		respBody, _ = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
	}
	if err != nil {
		logger.Error("Errored when sending request to the server", err)
		return SourcesT{}, false
	}
	var sourcesJSON SourcesT
	err = json.Unmarshal(respBody, &sourcesJSON)
	if err != nil {
		logger.Error("Errored while parsing request", err, string(respBody), resp.StatusCode)
		return SourcesT{}, false
	}
	return sourcesJSON, true
}