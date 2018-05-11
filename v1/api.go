package tracker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetProfile is used to request a profile from fortnite tracker
func GetProfile(platform, name, APIToken string) (Profile, error) {
	profile := Profile{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", ProfileRoute+platform+"/"+name, nil)
	if err != nil {
		log.Fatal("failed to create profile request", err)
		return profile, err
	}
	req.Header.Add(headerKeyName, APIToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("failed to get profile from server", err)
		return profile, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &profile)
	if err != nil {
		log.Fatal("failed to unmarshal profile json")
		return profile, err
	}
	return profile, nil
}
