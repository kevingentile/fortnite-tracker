// Copyright 2018 Kevin Gentile.
// Licensed under GNU General Public License v3.0
package tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

// Key is your API token
type Key struct {
	Value string
}

//LoadKey is used to load an API token from a file.
func LoadKey(path string) Key {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return Key{Value: string(data)}
}

func TestGetProfile(t *testing.T) {
	t.SkipNow()
	Key := LoadKey(testResourcePath + testKey)

	profile, err := GetProfile("pc", "laughingcabbage", Key.Value)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(profile)
}

func TestGetWins(t *testing.T) {
	profile, err := loadProfile()
	if err != nil {
		t.Error(err)
	}
	wins, err := GetWins(profile)
	if err != nil {
		t.Error(err)
	}
	if wins != 100 {
		t.Fail()
	}

}

func TestGetTop3(t *testing.T) {
	profile, err := loadProfile()
	if err != nil {
		t.Error(err)
	}
	top3, err := GetTop3s(profile)
	if err != nil {
		t.Error(err)
	}
	if top3 != 127 {
		t.Fail()
	}
}

func TestGetKills(t *testing.T) {
	profile, err := loadProfile()
	if err != nil {
		t.Error(err)
	}
	kills, err := GetKills(profile)
	if err != nil {
		t.Error(err)
	}
	if kills != 3961 {
		t.Fail()
	}
}

func TestGetKDR(t *testing.T) {
	profile, err := loadProfile()
	if err != nil {
		t.Error(err)
	}
	kdr, err := GetKDR(profile)
	if err != nil {
		t.Error(err)
	}
	if kdr != 1.46 {
		t.Fail()
	}
}

func TestGetCurrentKDR(t *testing.T) {
	profile, err := loadProfile()
	if err != nil {
		t.Error(err)
	}
	kdr, err := GetCurrentKDR(profile)
	if err != nil {
		t.Error(err)
	}
	if kdr != 4.16 {
		t.Fail()
	}
}

func loadProfile() (Profile, error) {
	profile := Profile{}
	data, err := ioutil.ReadFile(testResourcePath + "profile.json")
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(data, &profile)
	if err != nil {
		return profile, err
	}
	return profile, nil
}
