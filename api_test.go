package tracker

import (
	"fmt"
	"testing"

	"github.com/LaughingCabbage/tracker-bot/key"
)

func TestGetProfile(t *testing.T) {
	t.Skip()
	Key := key.LoadKey(testResourcePath + testKey)

	resp := GetProfile("pc", "laughingcabbage", Key.Value)
	//TODO validate response
	fmt.Println(string(resp))

}
