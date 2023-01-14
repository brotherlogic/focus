package main

import (
	"testing"

	"github.com/brotherlogic/godiscogs"
)

func TestEmptyImages(t *testing.T) {
	res := getImage([]*godiscogs.Image{})
	if res != "" {
		t.Errorf("Should have been blank: %v", res)
	}
}
