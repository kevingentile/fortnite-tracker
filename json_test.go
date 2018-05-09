package tracker

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestRecentMatch(t *testing.T) {

	actual := Match{
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

	data, err := ioutil.ReadFile(testResourcePath + "match.json")
	if err != nil {
		t.Error(err)
	}

	match := Match{}
	json.Unmarshal(data, &match)

	if !reflect.DeepEqual(match, actual) {
		t.Fail()
	}

}
