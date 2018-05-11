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
