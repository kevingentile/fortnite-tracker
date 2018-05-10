package tracker

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestStat(t *testing.T) {
	statField := actual := StatField{
		Label:        "TRN Rating",
		Field:        "TRNRating",
		Category:     "Rating",
		ValueInt:     906,
		Value:        "906",
		Percentile:   67,
		DisplayValue: "906",
	}
	actual := Stat{
		
	}
	data, err := ioutil.ReadFile(testResourcePath + "stat.json")
	if err != nil {
		t.Error(err)
	}
	stat := Stat{}
	json.Unmarshal(data, &stat)
	return

}

func TestStatField(t *testing.T) {
	actual := StatField{
		Label:        "TRN Rating",
		Field:        "TRNRating",
		Category:     "Rating",
		ValueInt:     906,
		Value:        "906",
		Percentile:   67,
		DisplayValue: "906",
	}
	data, err := ioutil.ReadFile(testResourcePath + "statField.json")
	if err != nil {
		t.Error(err)
	}
	statField := StatField{}
	json.Unmarshal(data, &statField)
	if !reflect.DeepEqual(actual, statField) {
		t.Fail()
	}
}
func TestLifeTimeKey(t *testing.T) {
	actual := LifeTimeKey{
		Key:   "Top 3",
		Value: "176",
	}
	data, err := ioutil.ReadFile(testResourcePath + "lifeTimeKey.json")
	if err != nil {
		t.Error(err)
	}
	lifeTimeKey := LifeTimeKey{}
	json.Unmarshal(data, &lifeTimeKey)
	if !reflect.DeepEqual(lifeTimeKey, actual) {
		t.Fail()
	}
}

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

// actual := LifeTimeStats{
// 	LifeTimeKey{
// 		Key:   "Top 3",
// 		Value: "176",
// 	},
// 	LifeTimeKey{
// 		Key:   "Top 5s",
// 		Value: "68",
// 	},
// 	LifeTimeKey{
// 		Key:   "Top 6s",
// 		Value: "218",
// 	},
// 	LifeTimeKey{
// 		Key:   "Top 12s",
// 		Value: "150",
// 	},
// 	LifeTimeKey{
// 		Key:   "Top 25s",
// 		Value: "368",
// 	},
// 	LifeTimeKey{
// 		Key:   "Score",
// 		Value: "486,054",
// 	},
// 	LifeTimeKey{
// 		Key:   "Matches Played",
// 		Value: "2810",
// 	},
// 	LifeTimeKey{
// 		Key:   "Wins",
// 		Value: "100",
// 	},
// 	LifeTimeKey{
// 		Key:   "Win%",
// 		Value: "4%",
// 	},
// 	LifeTimeKey{
// 		Key:   "Kills",
// 		Value: "3961",
// 	},
// 	LifeTimeKey{
// 		Key:   "K/d",
// 		Value: "1.46",
// 	},
// }
