// Copyright 2019 Kevin Gentile.
// Licensed under GNU General Public License v3.0

package tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

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
	wins, err := profile.GetWins()
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
	top3, err := profile.GetTop3s()
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
	kills, err := profile.GetKills()
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
	kdr, err := profile.GetKDR()
	if err != nil {
		t.Error(err)
	}
	if kdr != 1.46 {
		t.Fail()
	}
}

func TestLookupLifetimeStat(t *testing.T) {
	profile, err := loadProfile()
	if err != nil {
		t.Error(err)
	}

	statStr, err := lookupLifetimeStat(profile, "K/d")
	if err != nil {
		t.Error(err)
	}

	if statStr != "1.46" {
		t.Fail()
	}
}

// func TestGetCurrentKDR(t *testing.T) {
// 	profile, err := loadProfile()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	kdr, err := profile.GetCurrentKDR()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if kdr != 4.16 {
// 		t.Fail()
// 	}
// }

func loadProfile() (*Profile, error) {
	profile := &Profile{}
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
