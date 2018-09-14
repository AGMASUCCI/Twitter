package service_test

import (
	"github.com/Twitter/src/service"
	"testing"
)

//Los tests siempre arrancan con la palabra Test
func TestPublishedTweetIsSaved(t *testing.T) {

	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	if service.Tweet != tweet {
		t.Error("Expected tweet is", tweet)
	}

}
