package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "what-going-on.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"google.com",
		"what-going-on.com",
		"youtube.com",
	}

	want := map[string]bool{
		"google.com":        true,
		"what-going-on.com": false,
		"youtube.com":       true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
