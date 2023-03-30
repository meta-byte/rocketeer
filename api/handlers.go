package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/meta-byte/rocketeer-discord-bot/types"
)

func makeGETRequest(endpoint string, target interface{}) error {
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
	//read response into byte array and unmarshal into the target type.
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

func GetLaunchResults() (launch *types.LaunchResults, err error) {
	endpoint := os.Getenv("ENDPOINT")
	err = makeGETRequest(endpoint, &launch)
	if err != nil {
		return nil, err
	}
	if len(launch.Results) > 0 {
		return launch, nil
	}
	return nil, errors.New("no launch results")
}
