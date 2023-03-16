package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/meta-byte/rocketeer-discord-bot/types"
)

func makeHTTPRequest(endpoint string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Move business logic for getting the singular next launch to here. Setup separate function for multiple launches etc.
func GetLaunches() (launch *types.Launch, err error) {
	endpoint := os.Getenv("ENDPOINT")
	err = makeHTTPRequest(endpoint, &launch)
	if err != nil {
		return nil, err
	}
	if len(launch.Results) > 0 {
		return launch, nil
	}
	return nil, errors.New("no launch results")
}
