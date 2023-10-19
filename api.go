// Copyright 2019 Kevin Gentile.
// Licensed under GNU General Public License v3.0

package tracker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// ErrMissingToken returned when the fortnite tracker API token is not set
var ErrMissingToken = errors.New("Missing API token")

// GetProfile is used to request a profile from fortnite tracker
func GetProfile(platform, name, APIToken string) (*Profile, error) {
	if APIToken == "" {
		return nil, ErrMissingToken
	}
	profile := &Profile{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", ProfileRoute+platform+"/"+name, nil)
	if err != nil {
		log.Println("failed to create profile request", err)
		return profile, err
	}
	req.Header.Add(headerKeyName, APIToken)
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("failed to get profile from server", err)
		return profile, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response body")
	}

	err = json.Unmarshal(body, profile)
	if err != nil {
		log.Println("failed to unmarshal profile json")
		return profile, err
	}

	if profile.AccountID == "" {
		apiErr := Error{}
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, errors.New("error in profile request. API may be down")
		}
		return nil, errors.New("error in profile request: \"" + apiErr.Error + "\"")
	}

	return profile, nil
}

// GetWins returns a profile lifetime wins
func (profile *Profile) GetWins() (int, error) {
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

// GetTop3s returns a profile lifetime top 3 placements
func (profile *Profile) GetTop3s() (int, error) {
	top3String, err := lookupLifetimeStat(profile, "Top 3s")
	if err != nil {
		return -1, err
	}

	top3, err := strconv.Atoi(top3String)
	if err != nil {
		return -1, err
	}

	return top3, nil
}

// GetKills returns a profile lifetime kills
func (profile *Profile) GetKills() (int, error) {
	killsString, err := lookupLifetimeStat(profile, "Kills")
	if err != nil {
		return -1, err
	}

	kills, err := strconv.Atoi(killsString)
	if err != nil {
		return -1, err
	}

	return kills, nil
}

// GetKDR returns a profile lifetime kill death ratio
func (profile *Profile) GetKDR() (float64, error) {
	kdrString, err := lookupLifetimeStat(profile, "K/d")
	if err != nil {
		return -1.0, err
	}

	kdr, err := strconv.ParseFloat(kdrString, 64)
	if err != nil {
		return -1.0, err
	}
	return kdr, nil
}

func lookupLifetimeStat(profile *Profile, key string) (string, error) {
	for _, item := range profile.LifeTimeStats {
		if item.Key == key {
			return item.Value, nil
		}
	}
	return "-1.0", ErrNotFound
}

// DEPRECATED: Innacurate and requires alternate improvment. See GetKDR.
// GetCurrentKDR returns the kdr for the current season
// func (profile *Profile) GetCurrentKDR() (float64, error) {
// 	stats := profile.Stats

// 	totalKills := stats.CurrP10.Kills.ValueInt +
// 		stats.CurrP2.Kills.ValueInt +
// 		stats.CurrP9.Kills.ValueInt

// 	totalMatches := stats.CurrP10.Matches.ValueInt +
// 		stats.CurrP2.Matches.ValueInt +
// 		stats.CurrP9.Matches.ValueInt

// 	totalWins := stats.CurrP10.Top1.ValueInt +
// 		stats.CurrP2.Top1.ValueInt +
// 		stats.CurrP9.Top1.ValueInt

// 	kdrAverage := float64(totalKills) / float64(totalMatches-totalWins)

// 	kdr := math.Round(kdrAverage*100) / 100
// 	return kdr, nil
// }
