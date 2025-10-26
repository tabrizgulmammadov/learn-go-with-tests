package main

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "an url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
