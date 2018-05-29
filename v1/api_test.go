package tracker

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/LaughingCabbage/tracker-bot/key"
)

func TestGetProfile(t *testing.T) {
	t.SkipNow()
	Key := key.LoadKey(testResourcePath + testKey)

	_, err := GetProfile("pc", "laughingcabbage", Key.Value)
	if err != nil {
		t.Error(err)
	}
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
	if kdr != 4.53 {
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
