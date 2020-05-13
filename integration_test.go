package tracker

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestIntegration(t *testing.T) {
	enabledStr, ok := os.LookupEnv("FORTNITE_TRACKER_INTEGRATION")
	if !ok || enabledStr != "true" {
		fmt.Println("Skipping integration test")
		t.Skip()
	}
	profile, err := GetProfile("pc", "laughingcabbage", os.Getenv("FORTNITE_TRACKER_TOKEN"))
	if err != nil {
		t.Error(err)
	}

	_, err = profile.GetKills()
	if err != nil {
		t.Error(err)
	}

	_, err = profile.GetWins()
	if err != nil {
		t.Error(err)
	}

	statStr, err := lookupLifetimeStat(profile, "K/d")
	if err != nil {
		t.Error(err)
	}

	log.Println(statStr)
	statFloat, err := strconv.ParseFloat(statStr, 64)
	if err != nil {
		t.Error(err)
	}

	log.Println(statFloat)

	// kdr, err := profile.GetCurrentKDR()
	// if err != nil {
	// 	t.Error(err)
	// }

	// if math.IsNaN(kdr) {
	// 	t.Error()
	// }
}
