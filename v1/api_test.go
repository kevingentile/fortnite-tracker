package tracker

import (
	"testing"

	"github.com/LaughingCabbage/tracker-bot/key"
)

func TestGetProfile(t *testing.T) {
	t.Skip()
	Key := key.LoadKey(testResourcePath + testKey)

	_, err := GetProfile("pc", "laughingcabbage", Key.Value)
	if err != nil {
		t.Error(err)
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
