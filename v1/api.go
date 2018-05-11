package tracker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

// GetWins returns the number of wins from a users lifetime stats
func GetWins(profile Profile) (int, error) {
	winsString, err := lookupLifetimeStat(profile, "Wins")
	if err != nil {
		return -1, err
	}
	wins, err := strconv.Atoi(winsString)
	if err != nil {
		return -1, err
	}
	return wins, nil
}

	}
	return -1, errors.New("Wins not found in profile")
}

func lookupLifetimeStat(profile Profile, key string) (string, error) {
	for _, item := range profile.LifeTimeStats {
		if item.Key == key {
			return item.Value, nil
		}
	}
	return "", ErrNotFound
}
