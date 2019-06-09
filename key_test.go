// Copyright 2019 Kevin Gentile.
// Licensed under GNU General Public License v3.0

package tracker

import (
	"testing"
)

func TestLoadKey(t *testing.T) {
	//value taken directly from .test-tracker-key
	testKey := Key{Value: "1234123412341234ASDFasdf"}

	key := LoadKey(testResourcePath + "testkey")

	if *key != testKey {
		t.Fail()
	}
}
