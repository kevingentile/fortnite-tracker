package tracker

import (
	"fmt"
	"testing"

	"github.com/LaughingCabbage/tracker-bot/key"
)

func TestGetProfile(t *testing.T) {
	Key := key.LoadKey(testResourcePath + testKey)

	resp := GetProfile("PC", "laughingcabbage", Key.Value)
	//TODO validate response
	fmt.Println(resp)

}
