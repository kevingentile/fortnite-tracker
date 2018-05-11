package tracker

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestProfile(t *testing.T) {
	data, err := ioutil.ReadFile(testResourcePath + "profile.json")
	if err != nil {
		t.Error(err)
	}
	profile := Profile{}
	err = json.Unmarshal(data, &profile)
	if err != nil {
		t.Error(err)
	}
}

func TestStat(t *testing.T) {
	data, err := ioutil.ReadFile(testResourcePath + "stat.json")
	if err != nil {
		t.Error(err)
	}
	testStat := Stat{}
	err = json.Unmarshal(data, &testStat)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(testStat, stat) {
		t.Fail()
	}
}

func TestStatField(t *testing.T) {
	data, err := ioutil.ReadFile(testResourcePath + "statField.json")
	if err != nil {
		t.Error(err)
	}
	testStatField := StatField{}
	err = json.Unmarshal(data, &testStatField)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(testStatField, statField) {
		t.Fail()
	}
}
func TestLifeTimeKey(t *testing.T) {
	data, err := ioutil.ReadFile(testResourcePath + "lifeTimeKey.json")
	if err != nil {
		t.Error(err)
	}
	testLifeTimeKey := LifeTimeKey{}
	err = json.Unmarshal(data, &testLifeTimeKey)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(lifeTimeKey, testLifeTimeKey) {
		t.Fail()
	}
}

func TestRecentMatch(t *testing.T) {
	data, err := ioutil.ReadFile(testResourcePath + "match.json")
	if err != nil {
		t.Error(err)
	}

	testMatch := Match{}
	err = json.Unmarshal(data, &testMatch)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(testMatch, match) {
		t.Fail()
	}
}

var statField = StatField{
	Label:        "TRN Rating",
	Field:        "TRNRating",
	Category:     "Rating",
	ValueInt:     906,
	Value:        "906",
	Percentile:   67,
	DisplayValue: "906",
}

var stat = Stat{
	TRNRating:     statField,
	Score:         statField,
	Top1:          statField,
	Top3:          statField,
	Top5:          statField,
	Top6:          statField,
	Top10:         statField,
	Top12:         statField,
	Top25:         statField,
	KD:            statField,
	WinRatio:      statField,
	Matches:       statField,
	Kills:         statField,
	KPG:           statField,
	ScorePerMatch: statField,
}

var lifeTimeKey = LifeTimeKey{
	Key:   "Top 3",
	Value: "176",
}

var match = Match{
	ID:              201174559,
	AcccountID:      "e925c8a4-483a-412e-ad8f-32f474c743f5",
	Playlist:        "p10",
	Kills:           2,
	MinutesPlayed:   0,
	Top1:            0,
	Top5:            0,
	Top6:            0,
	Top10:           0,
	Top12:           1,
	Top25:           0,
	Matches:         1,
	Top3:            0,
	DateCollected:   "2018-05-07T00:10:01.72",
	Score:           330,
	Platform:        3,
	TRNRating:       1497.8,
	TRNRatingChange: 35.376,
}
